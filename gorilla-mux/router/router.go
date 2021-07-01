package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var routeList map[string]func(http.ResponseWriter, *http.Request) = {
	"/": Router.HomeHandler
}

type Router struct {
	router      *mux.Router
	Initialized bool
}



func (r *Router) Init() error {
	if r.initialized {
		return errors.New("Router already initialized")
	}
	r.router = mux.NewRouter()
	for k, v := range(routeList) {
		r.router.HandleFunc(k, v)
	}
	fmt.Println("Started the server")
	http.ListenAndServe(":3000", r.router)
	r.initialized = true
}

func (r * Router) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(w, "Hello")
}
