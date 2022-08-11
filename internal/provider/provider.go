package provider

import "github.com/sendgrid/sendgrid-go"

// Provider - a representation of the mailer provider.
// TODO: For now keeping it couple to the sendgrid client
//       But probably abstracting in later versions if other
//       providers are used
type Provider struct {
	Client *sendgrid.Client
}

// NewProvider - returns a new provider with sendgrid client
func NewProvider() *Provider {
	return &Provider{Client: sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))}
}
