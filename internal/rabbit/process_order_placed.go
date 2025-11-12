package rabbit

import (
	"time"

	"github.com/nmarsollier/cataloggo/internal/di"
	"github.com/nmarsollier/cataloggo/internal/env"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/rbt"
)

//	@Summary		Mensage Rabbit order_placed/order_placed
//	@Description	Cuando se recibe el mensage order_placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.
//	@Tags			Rabbit
//	@Accept			json
//	@Produce		json
//	@Param			order_placed	body	rschema.ConsumeOrderPlacedMessage	true	"Message order_placed"
//	@Router			/rabbit/order_placed [get]
//
// Consume Order Placed
func listenOrderPlaced(logger log.LogRusEntry) {
	for {
		err := rbt.ConsumeRabbitEvent[rschema.ConsumeOrderPlacedMessage](
			env.Get().FluentURL,
			env.Get().RabbitURL,
			env.Get().ServerName,
			"order_placed",
			"fanout",
			"catalog_order_placed",
			"",
			processOrderPlaced,
		)

		if err != nil {
			logger.Error(err)
		}
		logger.Info("RabbitMQ listenLogout conectando en 5 segundos.")
		time.Sleep(5 * time.Second)
	}
}

func processOrderPlaced(logger log.LogRusEntry, data *rschema.ConsumeOrderPlaced) {
	deps := di.NewInjector(logger)

	for _, a := range data.Message.Articles {
		art, err := deps.ArticleService().FindById(a.ArticleId)
		if err == nil {
			deps.ArticleService().DecrementStock(art.ID, a.Quantity)
		}
	}
}
