package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a struct which takes logger as one of it's element
type Hello struct{
	l *log.Logger
}

// NewHello is a function  return pointer Hello struct
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
	h.l.Println("Hello world")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw,"something went wrong",http.StatusBadRequest)
	}

	fmt.Fprintf(rw,"Hello %s\n",d)
}