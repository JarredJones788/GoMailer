package main

import (
	"config"
	"fmt"
	"mailer"
	"router"
)

func main() {

	config := &config.Config{
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

	mail := mailer.Emailer{}.Init(config)

	//Start router
	err := router.Router{}.Init(mail, config)

	if err != nil {
		fmt.Println(err)
		return
	}

}
