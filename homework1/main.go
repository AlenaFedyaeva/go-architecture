package main

import (
	"go-architecture/homework1/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {

	handler := &server{
		rep:repository.NewMapDB(),
	}
	router:= mux.NewRouter()
	router.HandleFunc("/hello",handler.hello).Methods("GET")
	setupServer(router, handler)
	

	srv:=&http.Server{
		Addr: ":8085",
		//Set timeouts to avoid Slowloris attacks
		// WriteTimeout: time.Second*10,
		// ReadTimeout: time.Second*15,
		// IdleTimeout: time.Second*60,
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}

func setupServer(router *mux.Router, handler *server) {
	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item", handler.listItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")
}
