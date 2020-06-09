package templates

import (
	"formats"
	"mailer"

	"gopkg.in/gomail.v2"
)

//TestTemplate - Structure of test email
type TestTemplate struct {
	recipients []mailer.Contact
}

//Create - return a test template
func (t TestTemplate) Create(recipients []mailer.Contact) *TestTemplate {
	t.recipients = recipients

	return &t
}

//GetRecipients - returns all contacts for this template
func (t *TestTemplate) GetRecipients() *[]mailer.Contact {
	return &t.recipients
}

//GetFirstRecipient - returns first contact in the list
func (t *TestTemplate) GetFirstRecipient() *mailer.Contact {
	return &t.recipients[0]
}

//RemoveFirstRecipient - removes first contact in the list
func (t *TestTemplate) RemoveFirstRecipient() bool {
	t.recipients = t.recipients[1:]
	return true
}

//SetupMessage - Setup the message that will be sent by this template
func (t *TestTemplate) SetupMessage(mainEmail string, domain string, contact *mailer.Contact) *gomail.Message {

	body := "Hello " + contact.Name
	html := formats.GetBasicFormat(body, "Test Header", domain)

	m := gomail.NewMessage()
	m.SetHeader("From", mainEmail)
	m.SetHeader("To", contact.Email)
	//m.SetAddressHeader("Cc", "jarred78888@hotmail.com", "Jarred")
	m.SetHeader("Subject", "Test Subject")
	m.SetBody("text/html", html)

	return m
}
