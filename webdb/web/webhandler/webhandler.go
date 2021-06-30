package webhandler

import (
	"fmt"
	"net/http"
	"strconv"
)

type WebHandler struct {
	Port       int16
	Started    bool
	ResetDB    bool
	PopulateDB bool
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World %s</h1>", r.URL.Path)
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
}

func (h *WebHandler) RegisterWeb() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
}

func (h *WebHandler) StartWeb() {
	fmt.Println("Server starting...")
	port := ":" + strconv.FormatInt(int64(h.Port), 10)
	http.ListenAndServe(port, nil)
}
