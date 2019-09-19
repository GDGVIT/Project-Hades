package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GDGVIT/Project-Hades/model"
	"github.com/GDGVIT/Project-Hades/organization/endpoints/views"
)

/**
* @api {get} /api/v1/org/join Send a join request to an org
* @apiName Send a join request to an org
* @apiGroup organization
* @apiPermission user
*
* @apiParam {string} org name of the organization as query param
*
*
* @apiParamExample {json} request-example
*
* curl -H '{"Authorization":"dbasjbdasbdbasdjkbasjkda"}' localhost/api/v1/org/join?org=GDG-VIT
*
* @apiParamExample {json} response-example
*{"message":"Some error occurred","data":"A join request is already pending"}
*
**/
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

/**
* @api {post} /api/v1/org/accept Accept join request
* @apiName Accept join request
* @apiGroup organization
* @apiParam admin
* @apiParam {string} org name of the organization as query param
*
*
* @apiParamExample {json} request-example
*{
*	"email":"test1@test.com",
*	"org":"GDG-VIT",
* "accept": true
*}
* @apiParamExample {json} response-example
*
*{"message":"Added as a member"}
*
**/

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

		access, err := model.EnforceRoleAdmin(tk.Email, req.Org)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
			return
		}

		if !access {
			json.NewEncoder(w).Encode(views.Msg{"failed to authenticate user", errors.New("failed to authenticate user")})
			return
		}

		if req.Accept {
			if err := model.AcceptJoinRequest(req.Email, req.Org); err != nil {
				json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
				return
			}
			if err := model.AddPolicy(req.Email, req.Org, "member"); err != nil {
				json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
				return
			}
		} else {
			if err := model.DenyJoinRequest(req.Email, req.Org); err != nil {
				json.NewEncoder(w).Encode(views.Msg{"Some error occurred", err.Error()})
				return
			}
		}

		json.NewEncoder(w).Encode(views.Msg{"Completed action", nil})
		return
	}
}
