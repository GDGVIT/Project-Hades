package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/GDGVIT/Project-Hades/model"
	"github.com/GDGVIT/Project-Hades/organization/endpoints/views"
)

// for users
func sendJoinRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tk, err := model.ValidateToken(token)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		org := r.URL.Query().Get("org")
		if err := model.CreateJoinRequest(tk.Email, org); err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		json.NewEncoder(w).Encode(views.Msg{"Join request send", nil})
		return
	}
}

// for admins
func acceptJoinRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tk, err := model.ValidateToken(token)
		if err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		req := views.AddMembers{}
		json.NewDecoder(r.Body).Decode(&req)
		if !model.Enforce(tk.Email, req.Org, "admin") {
			json.NewEncoder(w).Encode(views.Msg{"Error: User does not have sufficient permission", nil})
			return
		}
		if err := model.AcceptJoinRequest(req.Email, req.Org); err != nil {
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}

		if err := model.AddPolicy(req.Email, req.Org, "member"); err != nil {

			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}
		json.NewEncoder(w).Encode(views.Msg{"Added as a member", nil})
		return
	}
}
