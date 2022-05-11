package twiliosms

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"

	"packages/network"
)

var twilioAccountSid, twilioAuthTokent, url string

func InitTwilioSms(accountSid, authTokent string) {
	twilioAccountSid = accountSid
	twilioAuthTokent = authTokent
	url = "https://api.twilio.com/2010-04-01/Accounts/" + twilioAccountSid + "/Messages.json"
}

func SendSms(to, from, content string) (resp *http.Response, body string, err error) {
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", content)

	headers := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(twilioAccountSid+":"+twilioAuthTokent)),
		"Accept":        "application/json",
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	resp, body, err = network.SendRequest("POST", url, msgData.Encode(), headers)
	req, err := http.NewRequest("POST", url, &formData)

	return resp, body, err
}
