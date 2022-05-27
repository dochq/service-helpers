package twiliosms

import (
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/dochq/service-helpers/network"
)

var twilioAccountSid, twilioAuthTokent, twilioURL string

// InitTwilioSms - initialize needed variables
func InitTwilioSms(accountSid, authTokent, url string) {
	twilioAccountSid = accountSid
	twilioAuthTokent = authTokent
	twilioURL = url
}

// SendSms - send an sms via Twilio
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

	resp, body, err = network.SendRequest("POST", twilioURL, msgData.Encode(), headers)

	return resp, body, err
}
