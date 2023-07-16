package main

import (
	"regexp"
	"strings"

	"github.com/go-mail/mail"
)

var rxEmail = regexp.MustCompile(".+@.+\\\\..+")

type Message struct {
	Email   string
	content string
	Error   map[string]string
}

func (msg *Message) Validate() bool {
	msg.Error = make(map[string]string)

	matchEmail := rxEmail.Match([]byte(msg.Email))

	if matchEmail == false {
		msg.Error["Email"] = "Please enter valid Email!"
	}

	if strings.TrimSpace(msg.content) == "" {
		msg.Error["Content"] = "Please enter valid content!"
	}

	return len(msg.Error) == 0
}

func (msg *Message) Deliver() error {
	email := mail.NewMessage()
	email.SetHeader("To", "admin@example.com")
	email.SetHeader("From", "server@example.com")
	email.SetHeader("Reply-To", msg.Email)
	email.SetHeader("Subject", "New message via Contact Form")
	email.SetBody("text/plain", msg.content)

	username := "your_username"
	password := "your_password"

	return mail.NewDialer("smtp.mailtrap.io", 25, username, password).DialAndSend(email)
}
