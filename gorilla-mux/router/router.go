package router

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/geo-albin/projects/gorilla-mux/db"
	"github.com/geo-albin/projects/gorilla-mux/models"
	"github.com/gorilla/mux"
)

type Router struct {
	router      *mux.Router
	initialized bool
}

func (rt *Router) Init() error {
	if rt.initialized {
		return errors.New("Router already initialized")
	}
	rt.router = mux.NewRouter()
	rt.registerHandlers()
	fmt.Println("Started the server")
	http.ListenAndServe(":3000", rt.router)
	rt.initialized = true
	return nil
}

func (rt *Router) registerHandlers() error {
	if rt.initialized {
		return nil
	}
	rt.router.HandleFunc("/", rt.homeHandler)

	rt.router.HandleFunc("/product", rt.getProduct).Methods(http.MethodGet)
	rt.router.HandleFunc("/product", rt.putProduct).Methods(http.MethodPut)
	rt.router.HandleFunc("/product/{id:[0-9]+}", rt.postProduct).Methods(http.MethodPost)
	rt.router.HandleFunc("/product/{id:[0-9]+}", rt.deleteProduct).Methods(http.MethodDelete)
	return nil
}

func (rt *Router) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func (rt *Router) getProduct(w http.ResponseWriter, r *http.Request) {
	d := db.GetProducts()
	var p models.Products

	for _, data := range d {
		p = append(p, models.Product{
			ID:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
		})
	}
	if err := p.ToJSON(w); err != nil {
		fmt.Println(err.Error())
	}
}

func (rt *Router) putProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := p.FromJSON(r.Body); err != nil {
		fmt.Println(err.Error())
		return
	}
	p.ID = p.NextID()

	d := db.Data{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	}

	db.PutProduct(d)
}

func (rt *Router) postProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := p.FromJSON(r.Body); err != nil {
		fmt.Println(err.Error())
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}
	d := db.Data{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		UpdatedOn:   time.Now().UTC().String(),
	}

	db.PutProduct(d)
}

func (rt *Router) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	db.DeleteProduct(id)
}
