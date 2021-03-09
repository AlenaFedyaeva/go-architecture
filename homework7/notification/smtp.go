package notification

import (
	"fmt"
	"homework2/models"
	"net/smtp"
)

type smtpBot struct {
	smtp   int
	chatId int64
}

type SmtpNotification interface {
	MailOrder(order *models.Order) error
}

func NewSMPT(token string, chatID int64) (SmtpNotification, error) {
	return &smtpBot{
		chatId: chatID,
	}, nil
}

func (s *smtpBot) MailOrder(order *models.Order) error {

	// Sender data.
	from := "291407@gmail.com"
	password := "rjcnz29140Felix!"

	// Receiver email address.
	to := []string{
		"black_cat144@mail.ru",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"//"465"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")

	return nil
}

func tmp() {

}
