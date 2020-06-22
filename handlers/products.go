package handlers

import (
	"log"
	"net/http"

	"github.com/narenarjun/go-micro-learning/data"
)
// Products struct defines the product
type Products struct {
	l *log.Logger
}

// NewProducts func returns a pointer to Products
func  NewProducts(l *log.Logger) *Products  {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		p.getProducts(rw,r)
		return
	}

	// handle update request
	if r.Method == http.MethodPatch{}

	// catch all other calls
	rw.WriteHeader(http.StatusMethodNotAllowed)
}


func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request)  {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err!= nil {
		http.Error(rw, "Unable to marshal json",http.StatusInternalServerError)
		return
	}
}
