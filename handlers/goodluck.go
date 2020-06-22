package handlers

import (
	"log"
	"net/http"
)
// GoodLuck is a struct has a log as it's element
type GoodLuck struct{
	l *log.Logger
}

// NewGoodLuck is a function which returns pointer GoodLuck
func NewGoodLuck(l *log.Logger) *GoodLuck{
	return &GoodLuck{l}
}

func (g *GoodLuck) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	g.l.Println("hi ! GoodLuck")

	rw.Write([]byte("GoodLuck......!"))
}