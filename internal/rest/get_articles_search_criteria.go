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
//	@Success		200				{array}		article.ArticleData	"Articulos"
//	@Failure		400				{object}	errs.ValidationErr	"Bad Request"
//	@Failure		401				{object}	rst.ErrorData		"Unauthorized"
//	@Failure		404				{object}	rst.ErrorData		"Not Found"
//	@Failure		500				{object}	rst.ErrorData		"Internal Server Error"
//	@Router			/articles/search/:criteria [get]
//
// Obtener un articulo
func initGetArticlesSearchCriteria(engine *gin.Engine) {
	engine.GET(
		"/articles/search/:criteria",
		server.ValidateAuthentication,
		getArticleSearch,
	)
}

func getArticleSearch(c *gin.Context) {
	criteria := c.Param("criteria")

	deps := server.GinDi(c)

	result, err := deps.ArticleService().FindByCriteria(criteria)
	if err != nil {
		rst.AbortWithError(c, err)
		return
	}

	c.JSON(200, result)
}
