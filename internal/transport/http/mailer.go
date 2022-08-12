package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/anfelo/gomailer/internal/mailer"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
)

// MailerService - main interface for the mailer service
type MailerService interface {
	SendEmail(context.Context, mailer.EmailMessage) error
}

// Response - simple response struct
type Response struct {
	Message string
}

// SendEmailRequest - representation of the email request structure
type SendEmailRequest struct {
	From             string `json:"from" validate:"required"`
	FromName         string `json:"fromName"`
	To               string `json:"to" validate:"required"`
	ToName           string `json:"toName"`
	Subject          string `json:"subject" validate:"required"`
	PlainTextContent string `json:"plainTextContent" validate:"required"`
	HtmlContent      string `json:"htmlContent"`
}

// SendEmail - handles the request to send an email
func (h *Handler) SendEmail(w http.ResponseWriter, r *http.Request) {
	var emailRequest SendEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&emailRequest); err != nil {
		log.Error(err)
		http.Error(w, "not a valid email payload", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err := validate.Struct(emailRequest)
	if err != nil {
		log.Error(err)
		http.Error(w, "not a valid email payload", http.StatusBadRequest)
		return
	}

	emailMessage := mapSendEmailRequestToEmailMessage(emailRequest)
	err = h.Service.SendEmail(r.Context(), emailMessage)
	if err != nil {
		log.Error(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{"email sent successfully"}); err != nil {
		panic(err)
	}
}

func mapSendEmailRequestToEmailMessage(r SendEmailRequest) mailer.EmailMessage {
	return mailer.EmailMessage{
		From:             r.From,
		FromName:         r.FromName,
		To:               r.To,
		ToName:           r.ToName,
		Subject:          r.Subject,
		PlainTextContent: r.PlainTextContent,
		HtmlContent:      r.HtmlContent,
	}
}
