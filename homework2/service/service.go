package service

import (
	"errors"
	"homework2/models"
	"homework2/notification"
	"homework2/repository"
	"log"
)


type Service interface {
	CreateOrder(order *models.Order) (*models.Order, error)
}

type service struct {
	tg notification.Notification
	smtp notification.SmtpNotification 
	rep   repository.Repository
}

var (
	ErrItemNotExists = errors.New("item not exists")
)

func (s *service) CreateOrder(order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.rep.GetItem(itemID)
		if err != nil && err != repository.ErrNotFound {
			return nil, err
		}
		if err == repository.ErrNotFound {
			return nil, ErrItemNotExists
		}
	}

	order, err := s.rep.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	if err := s.tg.SendOrderCreated(order); err != nil {
		log.Println(err)
	}
	err= s.smtp.MailOrder(order)
	if  err != nil {
		log.Println(err)
	}
	return order, err
}

func NewService(rep repository.Repository, tg notification.Notification,smpt notification.SmtpNotification) Service {
	return &service{
		smtp: smpt,
		tg: tg,
		rep:   rep,
	}
}
