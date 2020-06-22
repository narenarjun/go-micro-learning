package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// ! creating a server
	s := &http.Server{
		Addr: ":8080",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func(){

	 err :=	s.ListenAndServe()
	 if err != nil{
		 l.Fatal(err)
	 }
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan,os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown",sig)

	// graceful shutdown -- is really important.
	/* 
	! this would allow us to close any opened database or other server functions or other leaks to close
	? before going into a abrubt shutdown
	*/
	// timoutcontext
	tc,_ := context.WithTimeout(context.Background(),30 * time.Second)
	 s.Shutdown(tc)

	// http.ListenAndServe(":8080",sm)
}