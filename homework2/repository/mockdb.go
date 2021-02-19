package repository

import (
	"homework2/models"
)

type mockRepository struct {
}

func NewMockDB() Repository {
	return &mockRepository{}
}
func (mock *mockRepository) ListItems(filter *ItemFilter) ([]*models.Item, error){
	return []*models.Item{
		&models.Item{
			ID:    1,
			Name:  "someName1",
			Price: 10,
		},
		&models.Item{
			ID:    2,
			Name:  "someName2",
			Price: 100,
		},
	} ,nil
}
func (mock *mockRepository) GetItem(itemId int32) (*models.Item, error){
	return &models.Item{
		ID:    1,
		Name:  "someName1",
		Price: 10,
	},nil
}
func (mock *mockRepository)  DeleteItem(itemId int32) error{
	return nil
}
func (mock *mockRepository)  UpdateItem(item *models.Item) (*models.Item, error){
	return &models.Item{
		ID:    1,
		Name:  "someUpdatetName1",
		Price: 10,
	},nil
}
func (mock *mockRepository) CreateItem(item *models.Item) (*models.Item, error){
	return &models.Item{
		ID:    1,
		Name:  "someUpdatetName1",
		Price: 10,
	},nil
}

func (mock *mockRepository) CreateOrder(item *models.Order) (*models.Order, error){
	return &models.Order{},nil
}
func (mock *mockRepository) ListOrders(filter *OrderFilter) ([]*models.Order, error){
	return []*models.Order{
		&models.Order{
			ID: 1,Phone: "123" ,Name: "1@mail.ru" ,
		},
		&models.Order{
			ID: 2,Phone: "23" ,Name: "2@mail.ru" ,
		},
	} ,nil
}
