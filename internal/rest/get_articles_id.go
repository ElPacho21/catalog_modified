package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/internal/rest/server"
	"github.com/nmarsollier/commongo/rst"
)

//	@Summary		Obtener un articulo
//	@Description	Obtener un articulo
//	@Tags			Catalogo
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer {token}"
//	@Param			articleId		path		string				true	"ID de articlo"
//	@Success		200				{object}	article.ArticleData	"Articulo"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/articles/:articleId [get]
//
// Obtener un articulo
func initGetArticles(engine *gin.Engine) {
	engine.GET(
		"/articles/:articleId",
		server.ValidateAuthentication,
		getArticle,
	)
}

func getArticle(c *gin.Context) {
	articleId := c.Param("articleId")

	deps := server.GinDi(c)
	result, err := deps.ArticleService().FindById(articleId)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, result)
}
