// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	endpoint "github.com/GDGVIT/Project-Hades/auth/pkg/endpoint"
	http "github.com/go-kit/kit/transport/http"
	http1 "net/http"
)

//  NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makeLoginHandler(m, endpoints, options["Login"])
	makeSignupHandler(m, endpoints, options["Signup"])
	makeCreateOrgHandler(m, endpoints, options["CreateOrg"])
	makeInviteHandler(m, endpoints, options["Invite"])
	makeShowInvitesHandler(m, endpoints, options["ShowInvites"])
	makeShowProfileHandler(m, endpoints, options["ShowProfile"])
	return m
}
