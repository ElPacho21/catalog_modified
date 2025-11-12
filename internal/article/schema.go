package article

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentMethod string

// Estuctura basica de del evento
type Article struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	Description Description        `bson:"description"  json:"description" validate:"required"`
	Price       float32            `bson:"price"  json:"price"`
	Stock       int                `bson:"stock"  json:"stock"`
	Created     time.Time          `bson:"created" json:"created"`
	Updated     time.Time          `bson:"updated" json:"updated"`
	Enabled     bool               `bson:"enabled" json:"enabled"`
}

type Description struct {
	Name        string `bson:"name"  json:"name" validate:"required,min=1,max=100"`
	Description string `bson:"description"  json:"description" validate:"required,min=1,max=256"`
	Image       string `bson:"image"  json:"image" validate:"max=100"`
}

// validateSchema valida la estructura para ser insertada en la db
func (e *Article) validateSchema() error {
	return validator.New().Struct(e)
}
