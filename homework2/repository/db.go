package repository

import (
	"errors"
	"fmt"
	"homework2/models"

	"sort"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

// Repository - intrface work with db
type Repository interface {
	CreateItem(item *models.Item) (*models.Item, error)
	UpdateItem(item *models.Item) (*models.Item, error)
	DeleteItem(itemId int32) error
	GetItem(itemId int32) (*models.Item, error)
	ListItems(filter *ItemFilter) ([]*models.Item, error)
	
	CreateOrder(order *models.Order) (*models.Order, error)
	ListOrders(filter *OrderFilter) ([]*models.Order, error)
}


type mapDB struct {
	mu         *sync.Mutex
	itemsTable *itemsTable
	ordersTable *ordersTable
}

type itemsTable struct {
	items map[int32]*models.Item
	maxID int32
}
type ordersTable struct {
	orders map[int32]*models.Order
	maxID  int32
}


func NewMapDB() Repository {
	return &mapDB{
		mu: &sync.RWMutex{},
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: &models.Item{
					ID:        1,
					Name:      "item1",
					Price:     50,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				2: &models.Item{
					ID:        2,
					Name:      "item2",
					Price:     60,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				3: &models.Item{
					ID:        3,
					Name:      "item3",
					Price:     70,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				4: &models.Item{
					ID:        4,
					Name:      "item4",
					Price:     80,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			maxID: 4,
		},
		ordersTable: &ordersTable{
			orders: make(map[int32]*models.Order),
			maxID:  0,
		},
	}
}

func (m *mapDB) CreateItem(item *models.Item) (*models.Item, error) {
	m.itemsTable.maxID++

	timeNow := time.Now().UTC()

	newItem := &models.Item{
		ID:        m.itemsTable.maxID,
		Price:     item.Price,
		Name:      item.Name,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	m.itemsTable.items[newItem.ID] = newItem

	return &models.Item{
		ID:        newItem.ID,
		Name:      newItem.Name,
		Price:     newItem.Price,
		CreatedAt: newItem.CreatedAt,
		UpdatedAt: newItem.UpdatedAt,
	}, nil
}

func (m *mapDB) ListItems(filter *ItemFilter) ([]*models.Item, error) {
	var res []*models.Item
	fmt.Println(m.itemsTable.items)
	//sort items
	itemSlice := make([]*models.Item, 0, len(m.itemsTable.items))
	for _, item := range m.itemsTable.items {
		itemSlice = append(itemSlice, item)
	}
	sort.Slice(itemSlice, func(i, j int) bool {
		return itemSlice[i].ID < itemSlice[j].ID
	})

	for _, item := range itemSlice {
		if filter.PriceLeft == nil && filter.PriceRight == nil {
			res = itemSlice
			break
		}
		fmt.Println(*filter.PriceLeft, item.Price, *filter.PriceRight)
		if (*filter.PriceLeft<=item.Price) && (item.Price <= *filter.PriceRight) {

			res = append(res, item)
		}
	}
	//If filter is off
	if filter.Limit==0 && filter.Offset==0 && filter.PriceLeft==nil && filter.PriceRight==nil {
		return res, nil
	}

	resFiltered := make([]*models.Item, 0, len(res))
	for idx, item := range res {
		if len(resFiltered) == filter.Limit {
			break
		}
		if idx < filter.Offset {
			continue
		}
		resFiltered = append(resFiltered, item)
	}
	return resFiltered, nil
}

func (m *mapDB) GetItem(ID int32) (*models.Item, error) {
	item, ok := m.itemsTable.items[ID]
	if !ok {
		return nil, ErrNotFound
	}

	return &models.Item{
		ID:        item.ID,
		Name:      item.Name,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}, nil
}

func (m *mapDB) DeleteItem(ID int32) error {
	_, ok := m.itemsTable.items[ID]
	if !ok {
		return ErrNotFound
	}

	delete(m.itemsTable.items, ID)
	return nil
}

func (m *mapDB) UpdateItem(item *models.Item) (*models.Item, error) {
	updateItem, ok := m.itemsTable.items[item.ID]
	if !ok {
		return nil, ErrNotFound
	}
	updateItem.Name = item.Name
	updateItem.Price = item.Price
	updateItem.UpdatedAt = time.Now().UTC()

	return &models.Item{
		ID:        updateItem.ID,
		Name:      updateItem.Name,
		Price:     updateItem.Price,
		CreatedAt: updateItem.CreatedAt,
		UpdatedAt: updateItem.UpdatedAt,
	}, nil
}


func (m *mapDB) CreateOrder(order *models.Order) (*models.Order, error) {
	m.ordersTable.maxID++

	timeNow := time.Now().UTC()

	newOrder := &models.Order{
		ID:            m.ordersTable.maxID,
		CustomerName:  order.CustomerName,
		CustomerPhone: order.CustomerPhone,
		ItemIDs:       order.ItemIDs,
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
	}

	m.mu.Lock()
	m.ordersTable.orders[newOrder.ID] = newOrder
	m.mu.Unlock()

	return &models.Order{
		ID:            newOrder.ID,
		CustomerName:  newOrder.CustomerName,
		CustomerPhone: newOrder.CustomerPhone,
		ItemIDs:       newOrder.ItemIDs,
		CreatedAt:     newOrder.CreatedAt,
		UpdatedAt:     newOrder.UpdatedAt,
	}, nil
}

func (m *mapDB) ListOrders(filter *OrderFilter) ([]*models.Order, error) {
	var res []*models.Order

	m.mu.RLock()
	orderSlice := make([]*models.Order, 0, len(m.ordersTable.orders))
	for _, order := range m.ordersTable.orders {
		orderSlice = append(orderSlice, order)
	}
	m.mu.RUnlock()

	sort.Slice(orderSlice, func(i, j int) bool {
		return orderSlice[i].ID < orderSlice[j].ID
	})

	resFiltered := orderSlice
	for idx, order := range res {
		if len(resFiltered) == filter.Limit {
			break
		}
		if idx < filter.Offset {
			continue
		}
		resFiltered = append(resFiltered, order)
	}
	return resFiltered, nil
}