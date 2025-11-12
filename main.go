package main

import (
	"github.com/nmarsollier/cataloggo/internal/di"
	"github.com/nmarsollier/cataloggo/internal/env"
	"github.com/nmarsollier/cataloggo/internal/graph"
	"github.com/nmarsollier/cataloggo/internal/rabbit"
	"github.com/nmarsollier/cataloggo/internal/rest"
	"github.com/nmarsollier/commongo/log"
)

//	@title			CatalogGo
//	@version		1.0
//	@description	Microservicio de Catalogo.
//	@contact.name	Nestor Marsollier
//	@contact.email	nmarsollier@gmail.com
//
//	@host			localhost:3002
//	@BasePath		/v1
//
// Main Method
func main() {
	deps := di.NewInjector(log.Get(env.Get().FluentURL, "cataloggo"))

	go graph.Start(deps.Logger())
	rabbit.Init(deps.Logger())
	rest.Start()
}
