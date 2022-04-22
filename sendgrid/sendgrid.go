package sendgrid

import (
	"encoding/base64"
	"fmt"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var sendgridClient *sendgrid.Client

type FileInfo struct {
	Name   string
	Type   string
	Buffer []byte
}

func InitSendgrid(key string) error {
	/*
		clientOptions := sentry.ClientOptions(sentry.ClientOptions{
			Dsn:              dsn,
			AttachStacktrace: true,
		})
		if err := sentry.Init(clientOptions); err != nil {
			return err
		}
	*/
	sendgridClient = sendgrid.NewSendClient(key)
	return nil
}

func SendGridEmail(sendGridEmailTmpl string, headers map[string]string, fromEmail *mail.Email, receipients map[string][]*mail.Email, subject string, dynamicTemplateData map[string]interface{}, files []*FileInfo) (*rest.Response, error) {
	var peopleToEmail, peoplceCcEmail, peoplceBccEmail []*mail.Email

	for _, receipient := range receipients["to"] {
		peopleToEmail = append(peopleToEmail, &mail.Email{
			Name:    receipient.Name,
			Address: receipient.Address,
		})
	}
	for _, receipient := range receipients["cc"] {
		peoplceCcEmail = append(peoplceCcEmail, &mail.Email{
			Name:    receipient.Name,
			Address: receipient.Address,
		})
	}
	for _, receipient := range receipients["bcc"] {
		peoplceBccEmail = append(peoplceBccEmail, &mail.Email{
			Name:    receipient.Name,
			Address: receipient.Address,
		})
	}

	sendData := &mail.SGMailV3{
		TemplateID: sendGridEmailTmpl,
		Subject:    subject,
		From: &mail.Email{
			Name:    fromEmail.Name,
			Address: fromEmail.Address,
		},
		Personalizations: []*mail.Personalization{
			{
				To:                  peopleToEmail,
				CC:                  peoplceCcEmail,
				BCC:                 peoplceBccEmail,
				DynamicTemplateData: dynamicTemplateData,
			},
		},
	}

	if headers != nil {
		sendData.Headers = headers
	}

	for _, file := range files {
		sendData.AddAttachment(&mail.Attachment{
			Content:     base64.StdEncoding.EncodeToString(file.Buffer),
			Type:        file.Type,
			Filename:    file.Name,
			Disposition: "attachment",
		})
	}

	res, err := sendgridClient.Send(sendData)
	if err != nil {
		return res, fmt.Errorf("sendgridClient.Send: %s", err)
	}

	if res.StatusCode != 202 {
		return res, fmt.Errorf("res.StatusCode: %d, %s", res.StatusCode, res.Body)
	}

	return res, nil
}
