package rest

import (
	"github.com/gin-gonic/gin"
	rbtpub "github.com/nmarsollier/cataloggo/internal/rabbit"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/nmarsollier/cataloggo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
)

//	@Summary		Eliminar Artículo
//	@Description	Eliminar Artículo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string	true	"Bearer {token}"
//	@Param			articleId		path	string	true	"ID de articlo"
//	@Success		200				"No Content"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/articles/:articleId [delete]
//
// Eliminar Artículo
func initDeleteArticles(engine *gin.Engine) {
	engine.DELETE(
		"/articles/:articleId",
		server.ValidateAuthentication,
		deleteArticle,
	)
}

func deleteArticle(c *gin.Context) {
	articleId := c.Param("articleId")
	deps := server.GinDi(c)

	err := deps.ArticleService().Disable(articleId)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	// Notify other services (e.g., Questions) that this article was deleted
	// Exchange: "article_deleted" (fanout)
	_ = rbtpub.PublishArticleDeleted(
		deps.Logger(),
		&rschema.ArticleDeletedMessage{ArticleId: articleId},
	)

	c.JSON(200, "")
}
