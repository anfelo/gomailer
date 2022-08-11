package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/anfelo/gomailer/internal/mailer"
	"github.com/go-playground/validator"
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
	FromName         string `json:"fromName" validate:"required"`
	To               string `json:"to" validate:"required"`
	ToName           string `json:"toName" validate:"required"`
	Subject          string `json:"subject" validate:"required"`
	PlainTextContent string `json:"plainTextContent" validate:"required"`
	HtmlContent      string `json:"htmlContent" validate:"required"`
}

// SendEmail - handles the request to send an email
func (h *Handler) SendEmail(w http.ResponseWriter, r *http.Request) {
	var emailRequest SendEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&emailRequest); err != nil {
		http.Error(w, "not a valid email payload", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err := validate.Struct(emailRequest)
	if err != nil {
		http.Error(w, "not a valid email payload", http.StatusBadRequest)
		return
	}

	emailMessage := mapSendEmailRequestToEmailMessage(emailRequest)
	err = h.Service.SendEmail(r.Context(), emailMessage)
	if err != nil {
		log.Print(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{"email sent successfully"}); err != nil {
		panic(err)
	}
}

func mapSendEmailRequestToEmailMessage(s SendEmailRequest) mailer.EmailMessage {
	return mailer.EmailMessage{}
}
