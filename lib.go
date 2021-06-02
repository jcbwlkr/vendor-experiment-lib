package lib

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AddHelp puts a helpful handler on the router.
func AddHelp(r *mux.Router) {
	r.Handle("/help", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Try turning it off and back on again"))
	}))
}
