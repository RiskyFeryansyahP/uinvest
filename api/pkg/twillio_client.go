package pkg

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// TwillioClient is client of twillio handle field for messaging
type TwillioClient struct {
	AccountID string
	AuthToken string
	APIURL    string
	From      string
	To        string
	Body      string
}

// NewTwillioClient is create new client for twillio
func NewTwillioClient(accountID, authToken string, twillioPhoneNumber string) (TwillioMessage, error) {
	if accountID == "" || authToken == "" || twillioPhoneNumber == "" {
		return nil, fmt.Errorf("account id or auth token or twillio phone number of twillio can't be empty")
	}

	apiURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountID)

	client := &TwillioClient{
		AccountID: accountID,
		AuthToken: authToken,
		APIURL:    apiURL,
		From:      twillioPhoneNumber,
	}

	return client, nil
}

// SendMessage is send message into specific number
func (t *TwillioClient) SendMessage(toUserNumber string, message string) (*http.Response, error) {
	client := &http.Client{}

	messageData := url.Values{}
	messageData.Set("To", toUserNumber)
	messageData.Set("From", t.From)
	messageData.Set("Body", message)

	msgDataReader := strings.NewReader(messageData.Encode())

	req, _ := http.NewRequest("POST", t.APIURL, msgDataReader)
	req.SetBasicAuth(t.AccountID, t.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
