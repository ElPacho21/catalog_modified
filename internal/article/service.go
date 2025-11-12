package article

import (
	"time"

	"github.com/nmarsollier/commongo/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleService interface {
	UpdateArticle(articleId string, articleData *UpdateArticleData) error
	FindById(id string, deps ...interface{}) (*ArticleData, error)
	FindByCriteria(criteria string) ([]*ArticleData, error)
	CreateArticle(articleData *UpdateArticleData, deps ...interface{}) (*ArticleData, error)
	Disable(articleId string) error
	DecrementStock(articleId primitive.ObjectID, amount int) error
}

func NewArticleService(log log.LogRusEntry, repository ArticleRepository) ArticleService {
	return &articleService{
		log:        log,
		repository: repository,
	}
}

type articleService struct {
	log        log.LogRusEntry
	repository ArticleRepository
}

func (s *articleService) UpdateArticle(articleId string, articleData *UpdateArticleData) error {
	err := s.repository.UpdateDescription(articleId, Description{
		Name:        articleData.Name,
		Description: articleData.Description,
		Image:       articleData.Image,
	})

	if err != nil {
		return err
	}

	err = s.repository.UpdateStock(articleId, articleData.Stock)

	if err != nil {
		return err
	}

	err = s.repository.UpdatePrice(articleId, articleData.Price)

	if err != nil {
		return err
	}

	return nil
}

func (s *articleService) FindById(id string, deps ...interface{}) (*ArticleData, error) {

	article, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return toArticleData(article), nil
}

func (s *articleService) FindByCriteria(criteria string) ([]*ArticleData, error) {
	articles, err := s.repository.FindByCriteria(criteria)
	if err != nil {
		return nil, err
	}

	result := []*ArticleData{}
	for _, a := range articles {
		result = append(result, toArticleData(a))
	}

	return result, nil
}

func (s *articleService) CreateArticle(articleData *UpdateArticleData, deps ...interface{}) (*ArticleData, error) {
	article, err := s.repository.Insert(&Article{
		ID: primitive.NewObjectID(),
		Description: Description{
			Name:        articleData.Name,
			Description: articleData.Description,
			Image:       articleData.Image,
		},
		Price:   articleData.Price,
		Stock:   articleData.Stock,
		Enabled: true,
		Created: time.Now(),
		Updated: time.Now(),
	})

	if err != nil {
		return nil, err
	}
	return toArticleData(article), nil
}

func (s *articleService) Disable(articleId string) error {
	return s.repository.Disable(articleId)
}

func (s *articleService) DecrementStock(articleId primitive.ObjectID, amount int) error {
	return s.repository.DecrementStock(articleId, amount)
}
