package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/GDGVIT/Project-Hades/model"
	"github.com/GDGVIT/Project-Hades/organization/endpoints/views"
)

func getOrgs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("org")
		orgs, err := model.GetOrgs(search)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err})
			return
		}
		json.NewEncoder(w).Encode(views.Msg{"Successful", orgs})
		return
	}
}
