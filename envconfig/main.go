package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/tobiaszheller/blog-examples/envconfig/sftp"
)

type config struct {
	Timeout            time.Duration `default:"5m" `
	Debug              bool          `required:"true"`
	SupportedCountries []string      `envconfig:"SUPPORTED_COUNTRIES"`
	SftpConfig
}

type SftpConfig sftp.Config

func main() {
	var cfg config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", cfg)
}
