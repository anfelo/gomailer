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
	//sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	return &Provider{Client: sendgrid.NewSendClient("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")}
}
