package main

import (
	"config"
	"fmt"
	"mailer"
	"os"
	"router"
)

func getConfig(mode string) *config.Config {

	if mode == "development" {
		return &config.Config{
			Email: config.EmailConfig{
				Email:       "info@tester788.info",
				Password:    "n!_e6b62ft$3fd",
				SMTPAddress: "smtp.office365.com",
				SMTPPort:    587,
			},
			EmailKey: "./dev_secrets/email_key.txt",
			Port:     ":5000",
			Host:     "DOMAIN",
		}
	}

	return &config.Config{
		Email: config.EmailConfig{
			Email:       "info@tester788.info",
			Password:    "n!_e6b62ft$3fd",
			SMTPAddress: "smtp.office365.com",
			SMTPPort:    587,
		},
		EmailKey: "/run/secrets/email_key",
		Port:     ":5000",
		Host:     "DOMAIN",
	}
}

func main() {

	//Default Dev config
	configType := "development"

	//Check if production flag is passed.
	if len(os.Args) > 1 {
		if os.Args[1] == "-production" {
			configType = "production"
		}
	}

	//Get correct runtime config
	config := getConfig(configType)

	mail := mailer.Emailer{}.Init(config)

	//Start router
	err := router.Router{}.Init(mail, config)

	if err != nil {
		fmt.Println(err)
		return
	}

}
