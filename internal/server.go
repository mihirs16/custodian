package internal

import (
	"net/http"

	"github.com/mihirs16/custodian/handlers"
	"github.com/mihirs16/custodian/models"
)

type Env struct {
	TypeDef models.TypeDefinitionModel
}

func SpawnServer(env *Env) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/types", handlers.MakeTypeDefHandler(&env.TypeDef))
	var httpHandler http.Handler = mux
	return httpHandler
}
