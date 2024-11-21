package main

import (
	"context"
	"fmt"

	pubsubrouter "github.com/sofyan48/pubsub-router"
	"github.com/sofyan48/pubsub-router/example/router"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// cfg := &pubsubrouter.Config{
	// 	Type:                    "service_account",
	// 	ProjectID:               os.Getenv("GOOGLE_PROJECT_ID"),
	// 	PrivateKeyID:            os.Getenv("GOOGLE_PRIVATE_KEY_ID"),
	// 	PrivateKey:              os.Getenv("GOOGLE_PRIVATE_KEY"),
	// 	ClientEmail:             os.Getenv("GOOGLE_CLIENT_EMAIL"),
	// 	ClientID:                os.Getenv("GOOGLE_CLIENT_ID"),
	// 	AuthURI:                 os.Getenv("GOOGLE_AUTH_URI"),
	// 	TokenURI:                os.Getenv("GOOGLE_TOKEN_URI"),
	// 	AuthProviderX509CertURL: os.Getenv("GOOGLE_AUTH_PROVIDER"),
	// 	ClientX509CertURL:       os.Getenv("GOOGLE_CLIENT_CERT_URL"),
	// }
	// sv := pubsubrouter.NewServer(context.Background(), cfg)
	sv := pubsubrouter.NewServerAutoConfig(context.Background(), "kirimin-aja")

	// publish data
	result, err := sv.Publish("cekaja", "/test", "Message send test")
	if err != nil {
		fmt.Println("error: ", err)
		panic(err)
	}
	fmt.Println("result Publish:> ", result)
	// // subscribe data
	rtr := router.NewRouter()

	sv.Subscribe("cekaja", rtr.Route()).Start()

}
