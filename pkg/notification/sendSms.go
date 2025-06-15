package notification

import (
	"encoding/json"
	"fmt"
	"go-ecommerce-app/config"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSMS(phoneNumber string, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{config: config}
}

// twilio
// SendSMS implements NotificationClient.
func (n *notificationClient) SendSMS(phoneNumber string, message string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: n.config.TwilioAccountSID,
		Password: n.config.TwilioAuthToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(n.config.TwilioFromPhone)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return err
	}

	response, _ := json.Marshal(*resp)
	fmt.Println("Response: " + string(response))
	return nil
}



