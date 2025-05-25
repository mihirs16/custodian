package internal

import (
	"net/http"

	"github.com/mihirs16/custodian/handlers"
	"github.com/mihirs16/custodian/models"
)

type Env struct {
	CatalogueModel models.CatalogueModel
	EntityModel    models.EntityModel
}

func SpawnServer(env *Env) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/catalogue", handlers.MakeCatalogueHandler(&env.CatalogueModel))
	mux.Handle("/entities", handlers.MakeEntitiesHandler(&env.EntityModel))
	var httpHandler http.Handler = mux
	return httpHandler
}
