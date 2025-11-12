package rabbit

import (
	"time"

	"github.com/nmarsollier/cataloggo/internal/di"
	"github.com/nmarsollier/cataloggo/internal/env"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/rbt"
)

//	@Summary		Mensage Rabbit article_exist/article_exist
//	@Description	Otros microservicios nos solicitan validar articulos en el catalogo.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			article_exist	body	rschema.ConsumeArticleExist	true	"Message para article_exist"
//	@Router			/rabbit/article_exist [get]
//
// Validar Artículos
func listenArticleExist(logger log.LogRusEntry) {
	for {
		err := rbt.ConsumeRabbitEvent[rschema.ConsumeArticleExistMessage](
			env.Get().FluentURL,
			env.Get().RabbitURL,
			env.Get().ServerName,
			"article_exist",
			"direct",
			"catalog_article_exist",
			"article_exist",
			processArticleExist,
		)

		if err != nil {
			logger.Error(err)
		}
		logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
		time.Sleep(5 * time.Second)
	}
}

func processArticleExist(logger log.LogRusEntry, data *rschema.ConsumeArticleExist) {
	deps := di.NewInjector(logger)

	article, err := deps.ArticleService().FindById(data.Message.ArticleId)
	if err != nil {
		deps.ArticleExistPublisher().PublishTo(
			data.Exchange,
			data.RoutingKey,
			&rschema.ArticleExistMessage{
				ArticleId:   data.Message.ArticleId,
				ReferenceId: data.Message.ReferenceId,
				Valid:       false,
			},
		)
		return
	}

	deps.ArticleExistPublisher().PublishTo(
		data.Exchange,
		data.RoutingKey,
		&rschema.ArticleExistMessage{
			ArticleId:   data.Message.ArticleId,
			ReferenceId: data.Message.ReferenceId,
			Stock:       article.Stock,
			Price:       article.Price,
			Valid:       article.Enabled,
		},
	)
}
