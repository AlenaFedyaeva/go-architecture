package service

import (
	"errors"
	"homework2/models"
	"homework2/notification"
	tg "homework2/pkg/tgbot"
	"homework2/repository"
	"log"
)
type Service interface {
	CreateItem(item *models.Item) (*models.Item, error)
	CreateOrder(order *models.Order) (*models.Order, error)
}

type service struct {
	tg notification.Notification
	db repository.Repository
}

var (
	ErrItemNotExists = errors.New("item not exists")
)

func (s *service) CreateOrder(order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.db.GetItem(itemID)
		if err != nil && err != repository.ErrNotFound {
			return nil, err
		}
		if err == repository.ErrNotFound {
			return nil, ErrItemNotExists
		}
	}

	order, err := s.db.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	if err := s.tg.SendOrderCreated(order); err != nil {
		log.Println(err)
	}
	return order, err
}


func NewService(tg tg.TelegramAPI, db repository.Repository) Service {
	return &service{
		db: db,
		tg: tg,
	}
}

func (s *service) CreateItem(item *models.Item) (*models.Item, error) {
	if item.Name == "" {
		return nil, errors.New("item name is empty")
	}
	if item.Price <= 0 {
		return nil, errors.New("item price should be positive")
	}

	return s.db.CreateItem(item)
}
