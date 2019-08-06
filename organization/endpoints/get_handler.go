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

func getJoinRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tk, err := model.ValidateToken(token)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		org := r.URL.Query().Get("org")
		if !model.Enforce(tk.Email, org, "admin") {
			json.NewEncoder(w).Encode(views.Msg{"Error: User does not have sufficient permission", nil})
			return
		}

		data, err := model.GetJoinRequests(org)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err})
			return
		}

		json.NewEncoder(w).Encode(views.Msg{
			Message: "Successful",
			Data:    data,
		})
		return
	}
}

func GetEventsAndOrgs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tk, err := model.ValidateToken(token)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		events, orgs, err := model.GetUserDetails(tk.Email)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err})
			return
		}

		json.NewEncoder(w).Encode(views.Msg{"Successful", map[string]interface{}{
			"organizations": orgs,
			"events":        events,
		}})
		return
	}
}
