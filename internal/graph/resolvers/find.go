package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/internal/graph/model"
	"github.com/nmarsollier/cataloggo/internal/graph/tools"
)

func GetArticle(ctx context.Context, articleId string) (*model.Article, error) {
	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	env := tools.GqlDi(ctx)

	result, err := env.ArticleService().FindById(articleId)
	if err != nil {
		return nil, err
	}

	return mapArticleDataToModel(result), nil
}
