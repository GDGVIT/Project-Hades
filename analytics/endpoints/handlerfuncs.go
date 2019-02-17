package endpoints

import "net/http"

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Analytics func"))
	}
}
