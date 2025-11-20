package resolvers

import (
	"context"

	"github.com/nmarsollier/cataloggo/internal/graph/tools"
	rbtpub "github.com/nmarsollier/cataloggo/internal/rabbit"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
)

func DeleteArticle(ctx context.Context, articleId string) (bool, error) {
	_, err := tools.ValidateAdmin(ctx)
	if err != nil {
		return false, err
	}

	env := tools.GqlDi(ctx)

	err = env.ArticleService().Disable(articleId)
	if err != nil {
		return false, err
	}

	_ = rbtpub.PublishArticleDeleted(
		env.Logger(),
		&rschema.ArticleDeletedMessage{ArticleId: articleId},
	)

	return true, nil
}
