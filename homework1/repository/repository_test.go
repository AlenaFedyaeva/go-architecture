package repository

import (
	"go-architecture/homework1/models"
	"testing"
)


func TestDeleteItem(t *testing.T) {
	db := NewMapDB()

	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	expected := &models.Item{
		ID:    1,
		Name:  input.Name,
		Price: input.Price,
	}
	// 1) Create item in DB
	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected result: expected %s, result %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected result: expected %d, result %d", expected.Price, result.Price)
	}
	if expected.ID != result.ID {
		t.Errorf("unexpected result: expected %d, result %d", expected.ID, result.ID)
	}
	// 2) Delete existing Item
	err = db.DeleteItem(result.ID)
	if err != nil {
		t.Error("unexpected err", err)
	}

	// 3) Try to Get deleted Item from DB
	result, err = db.GetItem(1)
	if err != ErrNotFound {
		t.Error("unexpected result: ", result.ID)
	}
	result, err = db.GetItem(100)
	if err != ErrNotFound {
		t.Error("unexpected result: ", result.ID)
	}
	//4) Try to delete wrong item
	// 2) Delete existing Item
	err = db.DeleteItem(100)
	if err != ErrNotFound {
		t.Error("unexpected err", err)
	}
}

func TestUpdateItem(t *testing.T) {
	db := NewMapDB()

	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	expected := &models.Item{
		ID:    1,
		Name:  input.Name,
		Price: input.Price,
	}
	// 1) Create item in DB
	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected result: expected %s, result %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected result: expected %d, result %d", expected.Price, result.Price)
	}
	if expected.ID != result.ID {
		t.Errorf("unexpected result: expected %d, result %d", expected.ID, result.ID)
	}
	// 2) Update item
	input = &models.Item{
		ID:    1,
		Name:  "someUpdatetName",
		Price: 100,
	}
	expected = &models.Item{
		ID:    input.ID,
		Name:  input.Name,
		Price: input.Price,
	}
	result, err = db.UpdateItem(input)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected result: expected %s, result %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected result: expected %d, result %d", expected.Price, result.Price)
	}
	if expected.ID != result.ID {
		t.Errorf("unexpected result: expected %d, result %d", expected.ID, result.ID)
	}

	// 3) Try to update Item with wrong ID
	input = &models.Item{
		ID:    100,
		Name:  "someUpdatetName",
		Price: 100,
	}
	result, err = db.UpdateItem(input)
	if err != ErrNotFound {
		t.Error("unexpected err", err)
	}

}

func TestCreateItem(t *testing.T) {
	db := NewMapDB()

	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	expected := &models.Item{
		ID:    1,
		Name:  input.Name,
		Price: input.Price,
	}
	// 1) Check if item created
	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected result: expected %s, result %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected result: expected %d, result %d", expected.Price, result.Price)
	}
	if expected.ID != result.ID {
		t.Errorf("unexpected result: expected %d, result %d", expected.ID, result.ID)
	}
	// 2) Check if item in db ==expected
	result, err = db.GetItem(expected.ID)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected result: expected %s, result %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected result: expected %d, result %d", expected.Price, result.Price)
	}
	if expected.ID != result.ID {
		t.Errorf("unexpected result: expected %d, result %d", expected.ID, result.ID)
	}
	//3) check increment in db
	input = &models.Item{
		Name:  "someName2",
		Price: 19,
	}
	expected = &models.Item{
		ID:    2,
		Name:  input.Name,
		Price: input.Price,
	}
	result, err = db.CreateItem(input)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected result: expected %s, result %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected result: expected %d, result %d", expected.Price, result.Price)
	}
	//check incremented id
	if expected.ID != result.ID {
		t.Errorf("unexpected result: expected %d, result %d", expected.ID, result.ID)
	}

}
