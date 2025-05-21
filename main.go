package main

import (
	"net/http"

	"github.com/mihirs16/custodian/internal"
	"github.com/mihirs16/custodian/models"
)

func main() {
	// setup database
	db := internal.SetupDBConn()

	// bootstrap dependencies
	env := &internal.Env{
		TypeDef: models.TypeDefinitionModel{DB: db},
	}

	// spawn server
	server := internal.SpawnServer(env)
	http.ListenAndServe(":3000", server)
}
