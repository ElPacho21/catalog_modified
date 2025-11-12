package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/internal/article"
	"github.com/nmarsollier/cataloggo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
)

//	@Summary		Actualizar Artículo
//	@Description	Actualizar Artículo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string						true	"Bearer {token}"
//	@Param			body			body		article.UpdateArticleData	true	"Informacion del articulo"
//	@Success		200				{object}	article.ArticleData			"Articulo"
//	@Failure		400				{object}	errs.ValidationErr			"Bad Request"
//	@Failure		401				{object}	rst.ErrorData				"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData				"Not Found"
//	@Failure		500				{object}	rst.ErrorData				"Internal Server Error"
//	@Router			/articles/:articleId [post]
//
// Actualizar Artículo
func initPostArticlesId(engine *gin.Engine) {
	engine.POST(
		"/articles/:articleId",
		server.ValidateAuthentication,
		updateArticle,
	)
}

func updateArticle(c *gin.Context) {
	body := article.UpdateArticleData{}
	if err := c.ShouldBindJSON(&body); err != nil {
		rst.AbortWithError(c, err)
		return
	}
	articleId := c.Param("articleId")

	deps := server.GinDi(c)
	err := deps.ArticleService().UpdateArticle(articleId, &body)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, "")
}
