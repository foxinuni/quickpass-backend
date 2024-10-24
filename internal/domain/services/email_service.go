package services

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type SendgridEmailServiceOptions interface {
	GetSendgridAPIKey() string
	GetSendgridEmail() string
}

type SendgridEmailService struct {
	options SendgridEmailServiceOptions
}

func NewSendgridEmailService(options SendgridEmailServiceOptions) EmailService {
	return &SendgridEmailService{
		options: options,
	}
}

func (s *SendgridEmailService) SendEmail(reciever string, subject string, body string) error {
	to := mail.NewEmail("cliente", reciever)
	from := mail.NewEmail("Quickpass", s.options.GetSendgridEmail())
	client := sendgrid.NewSendClient(s.options.GetSendgridAPIKey())

	// Create the email
	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.Subject = subject

	// Add the recipient
	personalization := mail.NewPersonalization()
	message.AddPersonalizations(personalization)
	personalization.AddTos(to)

	// Add the content
	message.AddContent(
		mail.NewContent("text/plain", body),
	)

	// Send the email
	if _, err := client.Send(message); err != nil {
		fmt.Print("error sending email")
		return err
	}
	fmt.Printf("to email sent %q", reciever)

	return nil
}
