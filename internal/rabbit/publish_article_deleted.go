package rabbit

import (
	"encoding/json"

	"github.com/nmarsollier/cataloggo/internal/env"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/nmarsollier/commongo/log"
	amqp "github.com/streadway/amqp"
)

// PublishArticleDeleted emits a fanout event to the durable exchange "article_deleted"
// with payload { type: "ARTICLE_DELETED", articleId }.
func PublishArticleDeleted(logger log.LogRusEntry, msg *rschema.ArticleDeletedMessage) error {
	conn, err := amqp.Dial(env.Get().RabbitURL)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Error(err)
		return err
	}
	defer ch.Close()

	// Ensure exchange exists with durable=true to match consumers
	if err := ch.ExchangeDeclare(
		"article_deleted",
		"fanout",
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,
	); err != nil {
		logger.Error(err)
		return err
	}

	body, _ := json.Marshal(msg)
	if err := ch.Publish(
		"article_deleted",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
