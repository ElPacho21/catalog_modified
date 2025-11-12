package article

import (
	"context"

	"github.com/nmarsollier/commongo/db"
	"github.com/nmarsollier/commongo/errs"
	"github.com/nmarsollier/commongo/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrID = errs.NewValidation().Add("id", "Invalid")

type ArticleRepository interface {
	FindByCriteria(criteria string) ([]*Article, error)
	FindById(articleId string) (*Article, error)
	Insert(article *Article) (*Article, error)
	Disable(articleId string) error
	UpdateDescription(articleId string, description Description) error
	UpdatePrice(articleId string, price float32) error
	UpdateStock(articleId string, stock int) error
	DecrementStock(articleId primitive.ObjectID, amount int) error
}

func NewArticleRepository(log log.LogRusEntry, collection db.Collection) ArticleRepository {
	return &articleRepository{
		log:        log,
		collection: collection,
	}
}

type articleRepository struct {
	log        log.LogRusEntry
	collection db.Collection
}

func (r *articleRepository) FindByCriteria(criteria string) ([]*Article, error) {
	filter := DBCriteriaFilter{
		Or: []map[string]DBCriteriaElement{
			{"description.name": {RegEx: criteria, Options: "i"}},
			{"description.description": {RegEx: criteria, Options: "i"}},
		},
	}

	cur, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		r.log.Error(err)

		return nil, err
	}
	defer cur.Close(context.Background())

	result := []*Article{}
	for cur.Next(context.Background()) {
		article := &Article{}
		if err := cur.Decode(article); err != nil {
			r.log.Error(err)

			return nil, err
		}
		result = append(result, article)
	}

	return result, nil
}

type DBCriteriaFilter struct {
	Or []map[string]DBCriteriaElement `bson:"$or"`
}

type DBCriteriaElement struct {
	RegEx   string `bson:"$regex"`
	Options string `bson:"$options"`
}

func (r *articleRepository) FindById(articleId string) (*Article, error) {
	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		r.log.Error(err)

		return nil, ErrID
	}

	article := &Article{}
	filter := DbIdFilter{ID: _id}

	if err = r.collection.FindOne(context.Background(), filter, article); err != nil {
		r.log.Error(err)

		return nil, err
	}

	return article, nil
}

func (r *articleRepository) Insert(article *Article) (*Article, error) {
	if err := article.validateSchema(); err != nil {
		r.log.Error(err)

		return nil, err
	}

	if _, err := r.collection.InsertOne(context.Background(), article); err != nil {
		r.log.Error(err)

		return nil, err
	}

	return article, nil
}

// disable Deshabilita el articulo para que no se pueda usar mas
func (r *articleRepository) Disable(articleId string) error {
	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		r.log.Error(err)

		return ErrID
	}

	_, err = r.collection.UpdateOne(context.Background(),
		DbIdFilter{ID: _id},
		DbEnableDocument{
			Set: DbEnableBody{
				Enabled: false,
			},
		},
		nil,
	)

	return err
}

type DbEnableDocument struct {
	Set DbEnableBody `bson:"$set"`
}

type DbEnableBody struct {
	Enabled bool `bson:"enabled" json:"enabled"`
}

// Actualiza la descripción de un articulo.
func (r *articleRepository) UpdateDescription(articleId string, description Description) error {
	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		r.log.Error(err)

		return ErrID
	}
	_, err = r.collection.UpdateOne(context.Background(),
		DbIdFilter{ID: _id},
		DbUpdateDescriptionDocument{
			Set: DbUpdateDescriptionBody{
				Description: description,
			},
		},
		nil,
	)

	return err
}

type DbUpdateDescriptionDocument struct {
	Set DbUpdateDescriptionBody `bson:"$set"`
}

type DbUpdateDescriptionBody struct {
	Description Description `bson:"description"  json:"description" validate:"required"`
}

// Actualiza el precio de un articulo.
func (r *articleRepository) UpdatePrice(articleId string, price float32) error {
	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		r.log.Error(err)

		return ErrID
	}
	_, err = r.collection.UpdateOne(context.Background(),
		DbIdFilter{ID: _id},
		DbUpdatePriceDocument{
			Set: DbUpdatePriceBody{
				Price: price,
			},
		},
		nil,
	)

	return err
}

type DbUpdatePriceDocument struct {
	Set DbUpdatePriceBody `bson:"$set"`
}

type DbUpdatePriceBody struct {
	Price float32 `bson:"price"  json:"price"`
}

// Actualiza el stock de un articulo.
func (r *articleRepository) UpdateStock(articleId string, stock int) error {
	_id, err := primitive.ObjectIDFromHex(articleId)
	if err != nil {
		r.log.Error(err)

		return ErrID
	}
	_, err = r.collection.UpdateOne(context.Background(),
		DbIdFilter{ID: _id},
		DbUpdateStockDocument{
			Set: DbUpdateStockBody{
				Stock: stock,
			},
		},
		nil,
	)
	if err != nil {
		r.log.Error(err)
	}

	return err
}

type DbUpdateStockDocument struct {
	Set DbUpdateStockBody `bson:"$set"`
}

func (r *articleRepository) DecrementStock(articleId primitive.ObjectID, amount int) error {
	_, err := r.collection.UpdateOne(context.Background(),
		bson.M{"_id": articleId},
		bson.M{
			"$inc": bson.M{
				"stock": -amount,
			}},
		nil,
	)

	if err != nil {
		r.log.Error(err)
	}

	return err
}

type DbIdFilter struct {
	ID primitive.ObjectID `bson:"_id"`
}

type DbUpdateStockBody struct {
	Stock int `bson:"stock"  json:"stock"`
}
