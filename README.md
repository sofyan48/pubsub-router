# GCP PubSub Router
[![Pubsub Router Local Development](https://github.com/sofyan48/pubsub-router/actions/workflows/docker-image.yml/badge.svg?branch=master)](https://github.com/sofyan48/pubsub-router/actions/workflows/docker-image.yml)
---
Route your action in gcp Pubsub easy

## Installing
```
go get github.com/sofyan48/pubsub-router
```

### Setup Client 
``` Golang
package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	pubsubrouter "github.com/sofyan48/pubsub-router"
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
}

```
### Setup Handler 
``` Golang

func handlerMessage() HandlerFunc {
	return func(m *pubsub.Message) error {
		fmt.Println("FROM EVENT:> ", string(m.Data))
		return nil
	}
}

func handlerMessage2() HandlerFunc {
	return func(m *pubsub.Message) error {
		fmt.Println("FROM TEST:> ", string(m.Data))
		return nil
	}
}
```

### Setup Router and Handler
``` Golang
	// add router
	rtr := pubsubrouter.NewRouter()
	
	// setup routing and handler
	rtr.Handle("/event", handlerMessage())
	rtr.Handle("/test", handlerMessage2())

```
### Starting Subscriber
``` Golang
	sv.Subscribe(os.Getenv("EVENT_BROKER_SERIAL"), rtr).Start()
```
### Starting Publisher 
Send event with attribute path
``` Golang
	sv.Publish(os.Getenv("EVENT_BROKER_SERIAL"), "/test", "Message send test")
```

## Example
For best model example check example folder

```
go run example/main.go
```