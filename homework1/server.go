package main

import (
	"encoding/json"
	"fmt"
	"go-architecture/homework1/models"
	"go-architecture/homework1/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type server struct {
	rep repository.Repository
}

func(s *server) hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func (s *server) createItemHandler(w http.ResponseWriter, r *http.Request) {
	item := new(models.Item)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// put item in db
	item, err := s.rep.CreateItem(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemId, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item := new(models.Item)
	item.ID=int32(itemId)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// put item in db
	item, err = s.rep.UpdateItem(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemId, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// del item in db
	err = s.rep.DeleteItem(int32(itemId))
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // возвращаем статус при успешном удалении
}

func (s *server) getItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get item in db
	item, err := s.rep.GetItem(int32(itemID))
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) parceItemFilterQuery(r *http.Request) *repository.ItemFilter{
	filter:=&repository.ItemFilter{}

	if limitRaw := r.FormValue("offset"); limitRaw != "" {
		if limit, err := strconv.Atoi(limitRaw); err== nil{
			filter.Limit=limit
		}	
	}
	if filter.Limit == 0 {
		filter.Limit = 5
	}
	if offsetRaw := r.FormValue("offset");  offsetRaw != "" {
		if offset, err := strconv.Atoi(offsetRaw); err== nil{
			filter.Offset=offset
		}
	}

	if priceRRaw := r.FormValue("price_right");  priceRRaw != "" {
		if priceR, err := strconv.ParseInt(priceRRaw,10,64); err== nil{
			filter.PriceRight=&priceR
		}
	}


	if priceLRaw := r.FormValue("price_left");  priceLRaw != "" {
		if priceL, err := strconv.ParseInt(priceLRaw,10,64); err== nil{
			filter.PriceLeft=&priceL
		}
	}
	return filter

}

type ListItemResponse struct{
	Payload []*models.Item `json: "payload"`
	Limit int `json: "limit"`
	Offset int `json: "offset"`
}

func (s *server) listItemHandler(w http.ResponseWriter, r *http.Request) {
	
	filter:= s.parceItemFilterQuery(r)

	// get item in db
	items, err := s.rep.ListItems(filter)
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp:= &ListItemResponse{Payload: items,Limit: filter.Limit,Offset: filter.Offset, }
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
