package main

import (
	"fmt"

	"github.com/anfelo/gomailer/internal/mailer"
	"github.com/anfelo/gomailer/internal/provider"
	transportHttp "github.com/anfelo/gomailer/internal/transport/http"
)

// Run - responsible for the instantiation
// and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	emailProvider := provider.NewProvider()
	mailerService := mailer.NewService(emailProvider)
	httpHandler := transportHttp.NewHandler(mailerService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Mailer API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
