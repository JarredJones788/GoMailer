package templates

import (
	"formats"
	"mailer"
)

//TestTemplate - Structure of test email
type TestTemplate struct {
	Recipients []mailer.Contact
}

//Create - return a test template
func (t TestTemplate) Create(recipients []mailer.Contact) *TestTemplate {
	t.Recipients = recipients

	return &t
}

//GetRecipients - returns all contacts for this template
func (t *TestTemplate) GetRecipients() *[]mailer.Contact {
	return &t.Recipients
}

//GetFirstRecipient - returns first contact in the list
func (t *TestTemplate) GetFirstRecipient() *mailer.Contact {
	return &t.Recipients[0]
}

//RemoveFirstRecipient - removes first contact in the list
func (t *TestTemplate) RemoveFirstRecipient() bool {
	t.Recipients = t.Recipients[1:]
	return true
}

//GetCC - returns a string of emails to be cc'd
func (t *TestTemplate) GetCC() string {
	return ""
}

//GetHeader - returns header of the email
func (t *TestTemplate) GetHeader() string {
	return "Test Header"
}

//GetSubject - returns subject of the email
func (t *TestTemplate) GetSubject() string {
	return "Test Subject"
}

//GetBody - returns body of the email
func (t *TestTemplate) GetBody(contact *mailer.Contact) string {
	body := "Hello " + contact.Name
	return formats.GetBasicFormat(body, t.GetHeader(), "")
}
