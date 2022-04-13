package sendgrid

import (
	"encoding/base64"
	"fmt"

	"github.com/getsentry/sentry-go"
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

func InitSendgrid(dsn, key string) error {
	clientOptions := sentry.ClientOptions(sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
	})
	if err := sentry.Init(clientOptions); err != nil {
		return err
	}
	sendgridClient = sendgrid.NewSendClient(key)
	return nil
}

func SendGridEmail(sendGridEmailTmpl string, fromEmail *mail.Email, receipients map[string][]*mail.Email, subject string, dynamicTemplateData map[string]interface{}, files []*FileInfo) (*rest.Response, error) {
	var peopleToEmail []*mail.Email

	for _, receipient := range receipients["to"] {
		peopleToEmail = append(peopleToEmail, &mail.Email{
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
				DynamicTemplateData: dynamicTemplateData,
			},
		},
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
		return res, fmt.Errorf("sendgridClient.Send: incorrect status code returned: %v, %s", res.StatusCode, res.Body)
	}

	return res, nil
}
