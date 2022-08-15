## GCP PubSub Router
Route your action in PubSub Easy

``` golang
package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	pubsubrouter "github.com/sofyan48/pubsub-router"
	"github.com/sofyan48/pubsub-router/handler"
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

	sv := pubsubrouter.NewServer(context.Background(), cfg)
	fmt.Println("RUN 1 WORKER TO ANY EVENT IN PUBSUB")
	rtr := pubsubrouter.NewRouter()
	rtr.Handle("/event", handlerMessage())
	rtr.Handle("/test", handlerMessage2())
	sv.Subscribe(os.Getenv("EVENT_BROKER_SERIAL"), rtr).Start()

}

func handlerMessage() handler.HandlerFunc {
	return func(m *pubsub.Message) error {
		fmt.Println("FROM EVENT:> ", string(m.Data))
		return nil
	}
}

func handlerMessage2() handler.HandlerFunc {
	return func(m *pubsub.Message) error {
		fmt.Println("FROM TEST:> ", string(m.Data))
		return nil
	}
}

```