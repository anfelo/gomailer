package http

import (
	"context"
	"net/http"

	"github.com/anfelo/gomailer/internal/mailer"
)

type MailerService interface {
	SendEmail(context.Context, mailer.EmailMessage) error
}

func (h *Handler) SendEmail(w http.ResponseWriter, r *http.Request) {
	//TODO: Implementation
}
