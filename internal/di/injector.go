package di

import (
	"github.com/nmarsollier/cataloggo/internal/article"
	"github.com/nmarsollier/cataloggo/internal/env"
	"github.com/nmarsollier/cataloggo/internal/rabbit/rschema"
	"github.com/nmarsollier/commongo/db"
	"github.com/nmarsollier/commongo/httpx"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/rbt"
	"github.com/nmarsollier/commongo/security"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

// Singletons
var database *mongo.Database
var httpClient httpx.HTTPClient
var catalogCollection db.Collection

type Injector interface {
	Logger() log.LogRusEntry
	Database() *mongo.Database
	HttpClient() httpx.HTTPClient
	SecurityRepository() security.SecurityRepository
	SecurityService() security.SecurityService
	CatalogCollection() db.Collection
	ArticleRepository() article.ArticleRepository
	ArticleService() article.ArticleService
	ArticleExistPublisher() rschema.ArticleExistPublisher
	ArticleDeletedPublisher() rschema.ArticleDeletedPublisher
}

type Deps struct {
	CurrLog               log.LogRusEntry
	CurrHttpClient        httpx.HTTPClient
	CurrDatabase          *mongo.Database
	CurrSecRepo           security.SecurityRepository
	CurrSecSvc            security.SecurityService
	CurrCatalogColl       db.Collection
	CurrArticleRepository article.ArticleRepository
	CurrArticleService    article.ArticleService
	CurrArtExistPublisher rschema.ArticleExistPublisher
	CurrArtDeletedPub     rschema.ArticleDeletedPublisher
}

func NewInjector(log log.LogRusEntry) Injector {
	return &Deps{
		CurrLog: log,
	}
}

func (i *Deps) Logger() log.LogRusEntry {
	return i.CurrLog
}

func (i *Deps) Database() *mongo.Database {
	if i.CurrDatabase != nil {
		return i.CurrDatabase
	}

	if database != nil {
		return database
	}

	database, err := db.NewDatabase(env.Get().MongoURL, "catalog")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}

	return database
}

func (i *Deps) HttpClient() httpx.HTTPClient {
	if i.CurrHttpClient != nil {
		return i.CurrHttpClient
	}

	if httpClient != nil {
		return httpClient
	}

	httpClient = httpx.Get()
	return httpClient
}

func (i *Deps) SecurityRepository() security.SecurityRepository {
	if i.CurrSecRepo != nil {
		return i.CurrSecRepo
	}
	i.CurrSecRepo = security.NewSecurityRepository(i.Logger(), i.HttpClient(), env.Get().SecurityServerURL)
	return i.CurrSecRepo
}

func (i *Deps) SecurityService() security.SecurityService {
	if i.CurrSecSvc != nil {
		return i.CurrSecSvc
	}
	i.CurrSecSvc = security.NewSecurityService(i.Logger(), i.SecurityRepository())
	return i.CurrSecSvc
}

func (i *Deps) CatalogCollection() db.Collection {
	if i.CurrCatalogColl != nil {
		return i.CurrCatalogColl
	}

	if catalogCollection != nil {
		return catalogCollection
	}

	userCollection, err := db.NewCollection(i.CurrLog, i.Database(), "catalog", IsDbTimeoutError)
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}

	return userCollection
}

func (i *Deps) ArticleRepository() article.ArticleRepository {
	if i.CurrArticleRepository != nil {
		return i.CurrArticleRepository
	}
	i.CurrArticleRepository = article.NewArticleRepository(i.Logger(), i.CatalogCollection())
	return i.CurrArticleRepository
}

func (i *Deps) ArticleService() article.ArticleService {
	if i.CurrArticleService != nil {
		return i.CurrArticleService
	}
	i.CurrArticleService = article.NewArticleService(i.Logger(), i.ArticleRepository())
	return i.CurrArticleService
}

// IsDbTimeoutError función a llamar cuando se produce un error de db
func IsDbTimeoutError(err error) {
	if err == topology.ErrServerSelectionTimeout {
		database = nil
		catalogCollection = nil
	}
}

func (i *Deps) ArticleExistPublisher() rschema.ArticleExistPublisher {
	if i.CurrArtExistPublisher != nil {
		return i.CurrArtExistPublisher
	}

	i.CurrArtExistPublisher, _ = rbt.NewRabbitPublisher[*rschema.ArticleExistMessage](
		rbt.RbtLogger(env.Get().FluentURL, "cataloggo", i.Logger().CorrelationId()),
		env.Get().RabbitURL,
		"",
		"direct",
		"",
	)

	return i.CurrArtExistPublisher
}

func (i *Deps) ArticleDeletedPublisher() rschema.ArticleDeletedPublisher {
	if i.CurrArtDeletedPub != nil {
		return i.CurrArtDeletedPub
	}

	// Pre-declare a fanout exchange used to broadcast article deletion events
	i.CurrArtDeletedPub, _ = rbt.NewRabbitPublisher[*rschema.ArticleDeletedMessage](
		rbt.RbtLogger(env.Get().FluentURL, "cataloggo", i.Logger().CorrelationId()),
		env.Get().RabbitURL,
		"article_deleted",
		"fanout",
		"",
	)

	return i.CurrArtDeletedPub
}
