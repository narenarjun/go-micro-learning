package main

import (
	"log"
	"net/http"
	"os"

	"github.com/narenarjun/go-micro-learning/handlers"
)

func main()  {
	
	l := log.New(os.Stdout,"product-api",log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodLuck(l)

	// ! creating a new servemux ourselves
	sm := http.NewServeMux()
	sm.Handle("/",hh)
	sm.Handle("/goodluck",gh)

	http.ListenAndServe(":8080",sm)
}