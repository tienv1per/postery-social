package mailer

import "github.com/sendgrid/sendgrid-go"

type MailTrapMailer struct {
	fromEmail string
	apiKey    string
	client    *sendgrid.Client
}
