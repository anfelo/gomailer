package mailer

import (
	"context"
	"fmt"
)

// EmailMessage - a representation of the email structure for our service this
// will contain all the fields needed to send the email
type EmailMessage struct {
	From             string
	FromName         string
	Subject          string
	To               string
	ToName           string
	PlainTextContent string
	HtmlContent      string
}

// Provider - the main interface that describes how to interact with the email
// provider
type Provider interface {
	SendEmail(context.Context, EmailMessage) error
}

// Service - is the struct on which all our logic will be built on top of
type Service struct {
	Provider Provider
}

// NewService - returns a pointer to a new service
func NewService(provider Provider) *Service {
	return &Service{
		Provider: provider,
	}
}

// SendEmail - sends an email with our email provider
func (s *Service) SendEmail(ctx context.Context, email EmailMessage) error {
	err := s.Provider.SendEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
