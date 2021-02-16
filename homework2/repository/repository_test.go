package repository

import (
	"go-architecture/homework1/models"
	"strconv"
	"testing"
)

func TestListItems(t *testing.T) {
	db := NewMapDB()

	expectedList := []*models.Item{}

	// 1) Create items in DB
	for i := 1; i <= 10; i++ {
		name := "SomeName_" + strconv.Itoa(i)
		input := &models.Item{
			Name:  name,
			Price: int64(i),
		}
		item, err := db.CreateItem(input)
		if err != nil {
			t.Error("unexpected err", err)
		}
		item.ID = int32(i)
		expectedList = append(expectedList, item)
	}
	// 2) get items with empty filter
	resultList, err := db.ListItems(&ItemFilter{})
	if err != nil {
		t.Error("unexpected err", err)
	}
	if len(resultList)!=len(expectedList){
		t.Error("unexpected result: wrong length of list")
	}

	// 3) get items with filter
	//Check if items equals in result & expected filtered List
	num1,num2:=3,8 //будем брать значения в интервале от [3,8] включая крайние значения

	left:=int64(num1)
	right:=int64(num2)

	expectedFilteredList:=expectedList[num1-1:num2]

	filter:=&ItemFilter{
		PriceLeft:  &left,
		PriceRight: &right,
		Limit:      6,
		Offset:     0,
	}
	resultFilteredList, err:= db.ListItems(filter)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if len(resultFilteredList)!=len(expectedFilteredList){
		t.Error("unexpected result: wrong length of filtered list")
	}
	for i, expected := range expectedFilteredList {
		
		if expected.Name != resultFilteredList[i].Name {
			t.Errorf("unexpected result: expected %s, result %s", expected.Name,
			resultFilteredList[i].Name)
		}
		if expected.Price != resultFilteredList[i].Price {
			t.Errorf("unexpected result: expected %d, result %d", expected.Price, 
			resultFilteredList[i].Price)
		}
		if expected.ID != resultFilteredList[i].ID {
			t.Errorf("unexpected result: expected %d, result %d", expected.ID,
			resultFilteredList[i].ID)
		}

	}
	//4) filter < list.size
	filter.Limit=3
	resultFilteredList, err= db.ListItems(filter)
	if err != nil {
		t.Error("unexpected err", err)
	}
	if len(resultFilteredList)!=filter.Limit{
		t.Errorf("unexpected result: limit %d result.len %d ", filter.Limit, len(resultFilteredList))
	}


}

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
