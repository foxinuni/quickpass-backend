package services

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioSMSServiceOptions interface {
	GetTwilioSID() string
	GetTwilioToken() string
	GetTwilioNumber() string
}

type SMSService interface {
	SendVerificationSMS(code string, phone string) error
}

type TwilioSMSService struct {
	options TwilioSMSServiceOptions
}

func NewTwilioSMSService(options TwilioSMSServiceOptions) SMSService {
	return &TwilioSMSService{
		options: options,
	}
}

func (t *TwilioSMSService) SendVerificationSMS(code string, phone string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: t.options.GetTwilioSID(),
		Password: t.options.GetTwilioToken(),
	})
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(t.options.GetTwilioNumber()) 
	params.SetBody("Your verification code is: " + code)

	_, err := client.Api.CreateMessage(params)
	return err
}