package main

import (
	"fmt"
	"go-architecture/homework1/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type Client struct {
	*http.Client
	URL string
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func executeRequest(req *http.Request,router *mux.Router) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    return rr
}

func Testhello(t *testing.T) {
	handler := &server{
		rep: repository.NewMockDB(),
	}

	router := mux.NewRouter() //a,router

	setupServer(router, handler)

	req, err := http.NewRequest("GET", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req,router)

	//checkResponseCode(t, http.StatusNotFound, response.Code)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
 
	fmt.Println(response.Body.String())

	// // Проверяем тело ответа
	// expected := `Parsed query-param with key "name": John`
	// if response.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		response.Body.String(), expected)
	// }
}
