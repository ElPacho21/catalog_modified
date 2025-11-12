package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nmarsollier/cataloggo/internal/env"
	"github.com/nmarsollier/cataloggo/internal/rest/server"
)

// Start this server
func Start() {
	engine := server.Router()
	InitRoutes(engine)
	engine.Run(fmt.Sprintf(":%d", env.Get().Port))
}

func InitRoutes(engine *gin.Engine) {
	initDeleteArticles(engine)
	initGetArticles(engine)
	initGetArticlesSearchCriteria(engine)
	initPostArticlesId(engine)
	initPostArticles(engine)
}
