package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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


// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("handle GET products")
	
	//  fetch the products from the datastore
	lp := data.GetProducts()

	// seralize thte list to JSON
	err := lp.ToJSON(rw)
	if err!= nil {
		http.Error(rw, "Unable to marshal json",http.StatusInternalServerError)
		return
	}
}

// AddProduct function adds a new product to the data store
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("handle Post products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

// UpdateProducts func updates a products values 
func (p *Products) UpdateProducts( rw http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	p.l.Println("handle PUT products", id)
	if err != nil {
		http.Error(rw,"Unable to convert ID",http.StatusBadRequest )
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	errdata := data.UpdateProduct(id , &prod)
	if errdata == data.ErrProductNotFound{
		http.Error(rw,"Product not found", http.StatusNotFound)
		return
	}

	if errdata != nil{
		http.Error(rw,"Product not found", http.StatusInternalServerError)
		return
	}
}
// KeyProduct is a key used for the context
type KeyProduct struct{}

// MiddleWareProductValidation is a middleware function 
func (p *Products) MiddleWareProductValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			// log.Fatalf("Error is: %v\n",err)
			p.l.Println("[ERROR] deserializing product",err)
			http.Error(rw,"Error reading product",http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product ",err)
			http.Error(
				rw, 
				fmt.Sprintf( "Error Validating Product: %s\n",err),
				http.StatusBadRequest)
			return
		}


		// add product to the context
		ctx := context.WithValue(r.Context(),KeyProduct{},prod)
		req := r.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw, req)
		} )
		
}

