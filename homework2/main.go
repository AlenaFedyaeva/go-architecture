package main

import (
	"flag"
	"fmt"
	"homework2/notification"
	"homework2/repository"
	"homework2/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var tokenTgStr string 
	var smtpStrLog,smtpStrPass, smtpReceiver string
	flag.StringVar(&tokenTgStr, "tg", "", "token for telegram api")
	flag.StringVar(&smtpStrLog, "smtpL", "", "log for smtp")
	flag.StringVar(&smtpStrPass, "smtpP", "", "pass for smtp")
	flag.StringVar(&smtpReceiver,"smtpReceiver","", "receiver smtp")
	flag.Parse()

	fmt.Println(tokenTgStr, smtpStrLog, smtpStrPass,smtpReceiver)
	tg, err := notification.NewTelegramBot(tokenTgStr, "*******")
	if err != nil {
		log.Fatal(err)
	}

	smtp,err:=notification.NewSMPT(smtpStrLog,smtpStrPass,smtpReceiver)

	rep := repository.NewMapDB()
	handler := &server{
		rep:     rep,
		service: service.NewService(rep, tg,smtp),
	}

	router := mux.NewRouter()
	setupServer(router, handler)

	srv := &http.Server{
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
	router.HandleFunc("/hello", handler.hello).Methods("GET")
	router.HandleFunc("/items", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/items", handler.listItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/items/{id}", handler.updateItemHandler).Methods("PUT")
	router.HandleFunc("/orders", handler.listOrdersHandler).Methods("GET")
	router.HandleFunc("/orders", handler.createOrderHandler).Methods("POST")
}
