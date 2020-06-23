package handlers

import (
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

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		// log.Fatalf("Error is: %v\n",err)
		http.Error(rw,"unable to unmarshal json",http.StatusBadRequest)
		return
	}

	// p.l.Printf("Prod: %v\n",prod) //!checked whether we got the products
	data.AddProduct(prod)
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

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		// log.Fatalf("Error is: %v\n",err)
		http.Error(rw,"unable to unmarshal json",http.StatusBadRequest)
		return
	}

	errdata := data.UpdateProduct(id ,prod)
	if errdata == data.ErrProductNotFound{
		http.Error(rw,"Product not found", http.StatusNotFound)
		return
	}

	if errdata != nil{
		http.Error(rw,"Product not found", http.StatusInternalServerError)
		return
	}
}
