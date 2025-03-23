package config

var Config = &config{}

type config struct {
	R2 R2
}

type R2 struct {
	AccessKey       string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`

	AccountID string `env:"ACCOUNT_ID"`
}
