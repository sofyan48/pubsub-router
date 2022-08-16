package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	pubsubrouter "github.com/sofyan48/pubsub-router"
	"github.com/sofyan48/pubsub-router/example/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error: ", err)
	}
	cfg := &pubsubrouter.Config{
		Type:                    "service_account",
		ProjectID:               os.Getenv("GOOGLE_PROJECT_ID"),
		PrivateKeyID:            os.Getenv("GOOGLE_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("GOOGLE_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("GOOGLE_CLIENT_EMAIL"),
		ClientID:                os.Getenv("GOOGLE_CLIENT_ID"),
		AuthURI:                 os.Getenv("GOOGLE_AUTH_URI"),
		TokenURI:                os.Getenv("GOOGLE_TOKEN_URI"),
		AuthProviderX509CertURL: os.Getenv("GOOGLE_AUTH_PROVIDER"),
		ClientX509CertURL:       os.Getenv("GOOGLE_CLIENT_CERT_URL"),
	}

	rtr := router.NewRouter()

	sv := pubsubrouter.NewServer(context.Background(), cfg)
	sv.Subscribe(os.Getenv("EVENT_BROKER_SERIAL"), rtr.Route()).Start()

}
