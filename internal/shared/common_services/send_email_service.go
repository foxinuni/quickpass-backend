package commonServices

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func SendEmail(receiverEmail string, subject string, content string) error {
	message := mail.NewV3Mail()
	from := mail.NewEmail(fmt.Sprintf("Quickpass"), os.Getenv("SENDGRID_EMAIL"))
	message.SetFrom(from)
	message.Subject = subject
	personalization := mail.NewPersonalization()
	to := mail.NewEmail("cliente", receiverEmail)
	personalization.AddTos(to)
	message.AddPersonalizations(personalization)
	message.AddContent(
		mail.NewContent("text/plain", content),
	)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}