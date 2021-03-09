package notification

import (
	"encoding/json"
	"fmt"
	"homework2/models"
	"net/smtp"
)

type smtpBot struct {
	login string `json:"login"`
	pass string `json:"pass"`
	receiver string `json:"receiver"`
}

type SmtpNotification interface {
	MailOrder(order *models.Order) error
}

func NewSMPT(login string, pass string, receiver string) (SmtpNotification, error) {
	return &smtpBot{
		login: login,
		pass: pass,
		receiver: receiver,
	}, nil
}

func (s *smtpBot) MailOrder(order *models.Order) error {

	// Sender data.
	from := s.login
	password := s.pass
	to:=[]string{s.receiver}

	
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := ":587"
	
	var message []byte
	str:=fmt.Sprintf(" order %d from: %s Phone: %s",order.ID,order.CustomerName,order.CustomerPhone,order.ItemIDs)
    message,err:=json.Marshal(str)
	// message := []byte("This is a test email message.")

	if err != nil {
		return err
	}
	fmt.Println(str)
	
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err= smtp.SendMail(smtpHost+smtpPort, auth, from,to , message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")

	return nil
}

func tmp() {

}
