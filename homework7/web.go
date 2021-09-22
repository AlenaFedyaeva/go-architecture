package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"homework2/models"
	"homework2/repository"
	"net/http"
)

var cli = &http.Client{}

type ShopClient interface {
	CreateItem(item *models.Item) (*models.Item, error)
	GetItem(req *GetItemsRequest) (*GetItemsResponse, error)
	BlockUser(req *BlockUserRequest) (*BlockUserResponse, error)
}

type GetItemsRequest struct {
	req *repository.ItemFilter
}

type GetItemsResponse struct {
	data []*models.Item
}

type CreateItemsRequest struct {
	name string
	price int64
}

type CreateItemsResponse struct {
	item models.Item
}

type BlockUserRequest struct {
	name string
	price int64
}

type BlockUserResponse struct {
	item models.Item
}


func BlockUser(request *BlockUserRequest) (*BlockUserResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	url:="http://localhost:8095/blockUser";
	req, err:= http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
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
		if resp.StatusCode > 399 {
			return nil, errors.New("something went wrong")
		}
	}
	if resp.StatusCode > 399 {
		return nil, errors.New("something went wrong")
	}
	// map - соответствие кодов http нашим ошибкам
	// if err := mapError(resp); err != nil {
		// return nil, err
	// }
	var response *BlockUserResponse
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil

}


func GetItem(request *GetItemsRequest) (*GetItemsResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	//Block - ItemsRequest in url 
	// url := fmt.Sprintf("http://localhost:8095/items/?limit=%d&offset=%d", request.Limit, request.Offset)
	// if *request.req.PriceRight > 0 {
		// url += fmt.Sprintf("&price_right=%d", *request.req.PriceRight)
	// }
	// if *request.req.PriceLeft > 0 {
		// url += fmt.Sprintf("&price_left=%d", *request.req.PriceLeft)
	// }
	url:="http://localhost:8095/getItems";
	req, err:= http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
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
		if resp.StatusCode > 399 {
			return nil, errors.New("something went wrong")
		}
	}
	if resp.StatusCode > 399 {
		return nil, errors.New("something went wrong")
	}
	// map - соответствие кодов http нашим ошибкам
	// if err := mapError(resp); err != nil {
		// return nil, err
	// }
	var response *GetItemsResponse // []*models.Item
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil

}

func CreateItem(request *CreateItemsRequest) (*CreateItemsResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	
	url:="http://localhost:8095/getItems";
	req, err:= http.NewRequest(http.MethodGet, url, bytes.NewReader(body))
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
		if resp.StatusCode > 399 {
			return nil, errors.New("something went wrong")
		}
	}
	if resp.StatusCode > 399 {
		return nil, errors.New("something went wrong")
	}
	
	var response *CreateItemsResponse // []*models.Item
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}
	return response, nil

}


func CreateItemOld(item *models.Item) (*models.Item, error) {
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
		if resp.StatusCode > 399 {
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
