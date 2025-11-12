package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/internal/graph/model"
	"github.com/nmarsollier/cataloggo/internal/graph/tools"
)

func FindArticleByID(ctx context.Context, articleId string) (*model.Article, error) {
	env := tools.GqlDi(ctx)

	result, err := env.ArticleService().FindById(articleId)
	if err != nil {
		return nil, err
	}

	return mapArticleDataToModel(result), nil
}
