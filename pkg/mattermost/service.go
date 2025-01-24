package mattermost

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/tech-djoin/wallet-djoin-service/internal/pkg/config"
)

func SendMessage(message string, fields []Field) {
	// Get environment variables
	webhookURL := config.Get("MATTERMOST_WEBHOOK_URL")
	username := config.Get("APP_NAME")

	// Generate message payload
	message = "**[PANIC]**" + " Some error occurred : " + message + "\n(ping @here)"

	// Generate attachments
	attachments := []map[string]interface{}{
		{
			"color":     "#FF0000",
			"mrkdwn_in": []string{"text"},
			"fields":    fields,
		},
	}

	// generate payload to send to mattermost
	payload := map[string]interface{}{
		"text":        message,
		"username":    username,
		"icon_url":    "https://wallet-test.solusisakti.xyz/images/logo/djoin-wallet.png",
		"attachments": attachments,
	}

	// Parse payload into json
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Post message into webhook
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(webhookURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Check if response is not successful
	if resp.StatusCode() != 200 {
		fmt.Printf("failed to send message to Mattermost: %s", resp.Status())
	}
}
