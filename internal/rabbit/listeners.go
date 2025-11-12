package rabbit

import (
	"github.com/nmarsollier/commongo/log"
)

func Init(logger log.LogRusEntry) {
	logger.
		WithField(log.LOG_FIELD_CONTROLLER, "Rabbit").
		WithField(log.LOG_FIELD_RABBIT_ACTION, "Init")

	go listenArticleExist(logger)

	go listenOrderPlaced(logger)

	go listenLogout(logger)
}
