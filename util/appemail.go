package util

import (
	"bytes"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)

type AppEmail interface {
	NewVerifyAccountEmail(username, email, verifyEmailToken string) *gomail.Message
}

type appEmail struct {
	verifyAccountTemplate *template.Template
	feVerificationURL     string
}

type AppEmailOpts struct {
	FEVerivicationURL string
}

func NewAppEmail(opts AppEmailOpts) (*appEmail, error) {
	verifyAccountTemplate, err := template.ParseFiles("templates/verify-email.html")
	if err != nil {
		return nil, err
	}

	return &appEmail{
		verifyAccountTemplate: verifyAccountTemplate,
		feVerificationURL:     opts.FEVerivicationURL,
	}, nil
}

func (a *appEmail) NewVerifyAccountEmail(username, email, verifyEmailToken string) *gomail.Message {
	var body bytes.Buffer
	a.verifyAccountTemplate.Execute(&body, struct {
		Username         string
		VerificationLink string
	}{
		Username:         username,
		VerificationLink: fmt.Sprintf("%s?email=%s&token=%s", a.feVerificationURL, email, verifyEmailToken),
	})
	mailer := gomail.NewMessage()
	mailer.SetHeader("Subject", "Welcome on Synapsis Online Store")
	mailer.SetBody("text/html", body.String())
	return mailer
}
