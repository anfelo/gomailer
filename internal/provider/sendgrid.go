package provider

import (
	"context"
	"errors"

	"github.com/anfelo/gomailer/internal/mailer"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	log "github.com/sirupsen/logrus"
)

var ErrBadRequest = errors.New("bad request")

// SendEmail - Sends a single email with sendgrid
func (p *Provider) SendEmail(ctx context.Context, emailData mailer.EmailMessage) error {
	from := mail.NewEmail(emailData.FromName, emailData.From)
	subject := emailData.Subject
	to := mail.NewEmail(emailData.ToName, emailData.To)
	plainTextContent := emailData.PlainTextContent
	htmlContent := emailData.HtmlContent

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := p.Client.Send(message)
	if err != nil {
		log.Error(err)
		return ErrBadRequest
	}

	if response.StatusCode >= 400 {
		log.Error(response)
		return ErrBadRequest
	}

	return nil
}
