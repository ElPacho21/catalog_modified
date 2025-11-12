package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/internal/graph/model"
	"github.com/nmarsollier/cataloggo/internal/graph/tools"
)

func SearchArticles(ctx context.Context, criteria string) ([]*model.Article, error) {

	_, err := tools.ValidateLoggedIn(ctx)
	if err != nil {
		return nil, err
	}

	env := tools.GqlDi(ctx)

	articles, err := env.ArticleService().FindByCriteria(criteria)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Article, len(articles))

	for i, a := range articles {
		result[i] = mapArticleDataToModel(a)
	}

	return result, nil
}
