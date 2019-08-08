package endpoints

import "net/http"

func Init() *http.ServeMux {
	mux := http.NewServeMux()

	// Auth endpoints
	mux.HandleFunc("/api/v1/org/create-org", createOrg())
	mux.HandleFunc("/api/v1/org/add-members", addMembers())
	mux.HandleFunc("/api/v1/org/login", login())
	mux.HandleFunc("/api/v1/org/login-org", loginOrg())
	mux.HandleFunc("/api/v1/org/signup", signup())

	// Get endpoints
	mux.HandleFunc("/api/v1/org/search", getOrgs())
	mux.HandleFunc("/api/v1/org/view-req", getJoinRequest())
	mux.HandleFunc("/api/v1/org/org-events", GetEventsAndOrgs())
	mux.HandleFunc("/api/v1/org/", GetOrgs())
	mux.HandleFunc("/api/v1/org/events", GetOrgEvents())

	// Join endpoints
	mux.HandleFunc("/api/v1/org/join", sendJoinRequest())
	mux.HandleFunc("/api/v1/org/accept", acceptJoinRequest())
	return mux
}
