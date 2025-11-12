package resolvers

import (
	"github.com/nmarsollier/cataloggo/internal/article"
	"github.com/nmarsollier/cataloggo/internal/graph/model"
)

func mapArticleDataToModel(article *article.ArticleData) *model.Article {
	return &model.Article{
		ID:          article.ID.Hex(),
		Name:        article.Name,
		Description: article.Description,
		Image:       article.Image,
		Price:       float64(article.Price),
		Stock:       article.Stock,
	}
}
