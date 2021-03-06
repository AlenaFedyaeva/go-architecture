package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"homework2/models"
	"homework2/repository"
	"net/http"
)

var cli = &http.Client{}

type ShopClient interface {
	CreateItem(item *models.Item) (*models.Item, error)
}

type GetItemResponse struct{
	data []*models.Item
	}// 

func GetItem(request *repository.ItemFilter )(*GetItemResponse, error) {
	body, err:=json.Marshal(request)
	if err!=nil{
		return nil, err
	}

	url:= fmt.Sprintf("http://localhost:8095/items/?limit=%d&offset=%d",request.Limit, request.Offset)

	if *req.PriceRight>0{
		url+= fmt.Sprintf("&price_right=%d",*req.PriceRight)
	}
	if *req.PriceLeft>0{
		url+= fmt.Sprintf("&price_left=%d",*req.PriceLeft)
	}

	req, err := http.NewRequest(http.MethodGet,url,bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 404:
		return nil, errors.New("item not found")
	default:
		if  resp.StatusCode > 399 {
			return nil, errors.New("something went wrong")
		}
	}
	if resp.StatusCode > 399 {
		return nil, errors.New("something went wrong")
	}

	if err:= mapError(resp); err!=nil {
		return nil,err
	}
	var response *GetItemResponse// []*models.Item
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil


}


func CreateItem(item *models.Item) (*models.Item, error) {
	body, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8095/items/", bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 404:
		return nil, errors.New("item not found")
	default:
		if  resp.StatusCode > 399 {
			return nil, errors.New("something went wrong")
		}
	}
	if resp.StatusCode > 399 {
		return nil, errors.New("something went wrong")
	}
	var resItem *models.Item
	if err := json.NewDecoder(resp.Body).Decode(resItem); err != nil {
		return nil, err
	}
	return resItem, nil

}
