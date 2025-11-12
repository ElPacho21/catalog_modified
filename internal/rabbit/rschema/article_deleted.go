package rschema

import "github.com/nmarsollier/commongo/rbt"

// ArticleDeletedPublisher publishes notifications when an article is deleted
// so other microservices (e.g., Questions) can react and cleanup related data.
type ArticleDeletedPublisher = rbt.RabbitPublisher[*ArticleDeletedMessage]

type ArticleDeletedMessage struct {	
	ArticleId string `json:"articleId"`
}
