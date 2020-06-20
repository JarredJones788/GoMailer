package types

import "mailer"

//EmailData - struct to send an email
type EmailData struct {
	Key      string           `json:"key"`
	Contacts []mailer.Contact `json:"contacts"`
	Template string           `json:"template"`
}
