package repository

import (
	"go-architecture/homework1/models"
	"testing"
)

func TestCreateItem(t *testing.T){
	db:=NewMapDB()

	input:=&models.Item{
		Name: "someName",
		Price: 10,
	}
	expected:=&models.Item{
		ID: 1,
		Name: input.Name,
		Price: input.Price,
	}
	// 1) Check if item created
	result,err:= db.CreateItem(input)
	if err!=nil {
		t.Error("unexpected err", err)
	}
	if expected.Name!= result.Name{
		t.Errorf("unexpected result: expected %s, result %s", expected.Name,result.Name)
	}
	if expected.Price!= result.Price{
		t.Errorf("unexpected result: expected %d, result %d", expected.Price,result.Price)
	}
	if expected.ID!= result.ID{
		t.Errorf("unexpected result: expected %d, result %d", expected.ID,result.ID)
	}
	// 2) Check if item in db ==expected
	result,err=db.GetItem(expected.ID)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if expected.Name!= result.Name{
		t.Errorf("unexpected result: expected %s, result %s", expected.Name,result.Name)
	}
	if expected.Price!= result.Price{
		t.Errorf("unexpected result: expected %d, result %d", expected.Price,result.Price)
	}
	if expected.ID!= result.ID{
		t.Errorf("unexpected result: expected %d, result %d", expected.ID,result.ID)
	}
	//3) check increment in db
	input=&models.Item{
		Name: "someName2",
		Price: 19,
	}
	expected=&models.Item{
		ID: 2,
		Name: input.Name,
		Price: input.Price,
	}
	result,err= db.CreateItem(input)
	if err!=nil {
		t.Error("unexpected err", err)
	}
	if expected.Name!= result.Name{
		t.Errorf("unexpected result: expected %s, result %s", expected.Name,result.Name)
	}
	if expected.Price!= result.Price{
		t.Errorf("unexpected result: expected %d, result %d", expected.Price,result.Price)
	}
	//check incremented id
	if expected.ID!= result.ID{
		t.Errorf("unexpected result: expected %d, result %d", expected.ID,result.ID)
	}

}