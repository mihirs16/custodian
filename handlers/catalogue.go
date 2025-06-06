package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mihirs16/custodian/models"
)

func MakeCatalogueHandler(t *models.CatalogueModel) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				hasType := r.URL.Query().Has("type")
				if !hasType {
					http.Error(w, "Missing `type` in query", http.StatusBadRequest)
					return
				}

				Type := r.URL.Query().Get("type")
				typeDef, err := t.FetchType(Type)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err = json.NewEncoder(w).Encode(&typeDef)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				return
			}

			if r.Method == http.MethodPost {
				var field models.FieldDefinitionOpts
				err := json.NewDecoder(r.Body).Decode(&field)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				err = t.CreateField(field)
				if err != nil {
					switch err.(type) {
					case *models.FieldDataTypeError:
						http.Error(w, err.Error(), http.StatusBadRequest)
					default:
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
					return
				}

				w.Write([]byte("Field created succesfully"))
				return
			}

			w.WriteHeader(http.StatusBadRequest)
		},
	)
}
