package servicefactory

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

// SendEmail send email
type SendEmail interface {
	FireEmail() (string, error)
}

// Email email struct
type Email struct {
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	ToEmail   string `json:"to_email"`
	ToName    string `json:"to_name"`
	Subject   string `json:"subject"`
	TextPart  string `json:"text_part"`
	HTMLPart  string `json:"html_part"`
	CustomID  string `json:"custom_id"`
}

// EmailJet email jet
type EmailJet struct {
	Email
}

// FireEmail fire email
func (e *EmailJet) FireEmail() (string, error) {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("KEY1"), os.Getenv("KEY2"))
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: e.FromEmail,
				Name:  e.FromName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: e.ToEmail,
					Name:  e.ToName,
				},
			},
			Subject:  e.Subject,
			TextPart: e.TextPart,
			HTMLPart: e.HTMLPart,
			CustomID: e.CustomID,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Printf("Data: %+v\n", res)
	return `{"msg": "Sent"}`, nil
}

// AwsEmail aws email
type AwsEmail struct {
	Email
}

// FireEmail fire email
func (e *AwsEmail) FireEmail() (string, error) {
	return "", nil
}

// GetServices service factory
func GetServices(e string, r *http.Request) (SendEmail, error) {
	var se SendEmail
	var em Email
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&em)
	if err != nil {
		fmt.Println(err)
		return se, errors.New("err")
	}
	if e == "emailjet" {
		return &EmailJet{Email: Email(em)}, nil
	} else if e == "awsemail" {
		return &AwsEmail{Email: Email(em)}, nil
	}
	return se, nil
}
