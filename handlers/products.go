package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	// handle post request
	if r.Method == http.MethodPost{
		p.addProduct(rw,r)
		return
	}

	if r.Method == http.MethodPut{
		// expect the id in the URI
		p.l.Println("handle Put method")
		regx := regexp.MustCompile(`/([0-9]+)`)
		 g := regx.FindAllStringSubmatch(r.URL.Path,-1)

		 if len(g) != 1{
			 p.l.Println("Invalid URL more than one id")
			 http.Error(rw,"Invalid URI", http.StatusBadRequest)
			 return
		 }

		 if len(g[0]) != 2{
			p.l.Println("Invalid URL more than one capture grooup")
			http.Error(rw,"Invalid URI", http.StatusBadRequest)
			return
		 }
		 idString := g[0][1]
		 id,err := strconv.Atoi(idString)
		 if err != nil{
			p.l.Println("Invalid URLunable to convert to number ", idString) 
			http.Error(rw,"Invalid URI", http.StatusBadRequest)
			return
		 }

		//  p.l.Println("got id:",id)
		p.updateProducts(id,rw,r)
		 return
	}


	// catch all other calls
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}


func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("handle GET products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err!= nil {
		http.Error(rw, "Unable to marshal json",http.StatusInternalServerError)
		return
	}
}


func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request)  {
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


func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("handle PUT products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		// log.Fatalf("Error is: %v\n",err)
		http.Error(rw,"unable to unmarshal json",http.StatusBadRequest)
		return
	}

	errdata := data.UpdateProduct(id,prod)
	if errdata == data.ErrProductNotFound{
		http.Error(rw,"Product not found", http.StatusNotFound)
		return
	}

	if errdata != nil{
		http.Error(rw,"Product not found", http.StatusInternalServerError)
		return
	}
}
