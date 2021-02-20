package main

import (
	"go-architecture/homework2/service"
	"homework2/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	// var tokenStr string
	// flag.StringVar(&tokenStr, "t", "", "token for telegram api")
	// flag.Parse()

	// tg, err := notification.NewTelegramBot(tokenStr, 323615875)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rep := repository.NewMapDB()
	handler := &server{
		rep:rep,
	  service: service.NewService(rep, nil),
	}

	router:= mux.NewRouter()
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
	router.HandleFunc("/hello",handler.hello).Methods("GET")
	router.HandleFunc("/items", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/items", handler.listItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/items/{id}", handler.updateItemHandler).Methods("PUT")
	router.HandleFunc("/orders",  handler.listOrdersHandler).Methods("GET")
	router.HandleFunc("/orders",  handler.createOrderHandler).Methods("POST")
}
