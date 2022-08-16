package pubsubrouter

import (
	"context"

	"github.com/sofyan48/pubsub-router/pkg/session"
)

type Config struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

func NewServer(ctx context.Context, cfg *Config) *Server {
	sess := session.New(ctx, &session.Config{
		Type:                    cfg.Type,
		ProjectID:               cfg.ProjectID,
		PrivateKeyID:            cfg.PrivateKeyID,
		PrivateKey:              cfg.PrivateKey,
		ClientEmail:             cfg.ClientEmail,
		ClientID:                cfg.ClientID,
		AuthURI:                 cfg.AuthURI,
		TokenURI:                cfg.TokenURI,
		AuthProviderX509CertURL: cfg.AuthProviderX509CertURL,
		ClientX509CertURL:       cfg.ClientX509CertURL,
	})
	return NewSession(ctx, sess)
}
