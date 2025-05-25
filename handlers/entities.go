package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mihirs16/custodian/models"
)

func MakeEntitiesHandler(c *models.EntityModel) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				hasId := r.URL.Query().Has("id")
				if !hasId {
					http.Error(w, "Missing `id` in query", http.StatusBadRequest)
					return
				}

				Entity := r.URL.Query().Get("id")
				entity, err := c.FetchEntity(Entity)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err = json.NewEncoder(w).Encode(&entity)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				return
			}

			if r.Method == http.MethodPost {
				var entityOpts models.EntityOptions
				err := json.NewDecoder(r.Body).Decode(&entityOpts)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				entityId, err := c.CreateEntity(entityOpts)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				fmt.Fprintf(w, "Entity %s created succesfully", entityId)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
		},
	)
}
