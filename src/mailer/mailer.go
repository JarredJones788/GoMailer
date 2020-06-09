package mailer

import (
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
)

//Contact - Structure of a contact
type Contact struct {
	Name  string
	Email string
}

//Template - interface
type Template interface {
	GetRecipients() *[]Contact
	GetFirstRecipient() *Contact
	RemoveFirstRecipient() bool
	SetupMessage(mainEmail string, domain string, contact *Contact) *gomail.Message
}

//Emailer - emailer struct
type Emailer struct {
	Email       string
	Password    string
	SMTPAddress string
	SMTPPort    int
	Host        string
	InProgress  bool
}

//Init - start email service
func (e Emailer) Init() *Emailer {
	e.Email = "info@tester788.info"
	e.Password = "n!_e6b62ft$3fd"
	e.SMTPAddress = "smtp.office365.com"
	e.SMTPPort = 587
	e.Host = "localhost:3000"
	e.InProgress = false
	return &e
}

//SendEmail - Sends template to the given emails inside the template
func (e *Emailer) SendEmail(template Template) error {

	//Create SMTP Dial
	d := gomail.NewDialer(e.SMTPAddress, e.SMTPPort, e.Email, e.Password)

	//Attempt to connect
	t, err := d.Dial()

	if err != nil {
		return err
	}

	defer t.Close() //Close connection when done using

	for {

		//If emails are already being sent then wait untill we can send again
		if e.InProgress {
			continue
		}

		//Set InProgress to true, no other emails can be sent until its reset to false
		e.InProgress = true

		//Get first contact in list
		contact := template.GetFirstRecipient()

		//Validate contact
		if contact == nil || contact.Email == "" {
			template.RemoveFirstRecipient() //If not valid remove and continue to next contact
			continue
		}

		//Remove first contact from the list
		template.RemoveFirstRecipient()

		//Create the email for the specific person
		m := template.SetupMessage(e.Email, e.Host, contact)

		//Attempt to send email
		if err := t.Send(e.Email, []string{contact.Email}, m); err != nil {
			//Log Email failed sending
			fmt.Println("Failed Sending Email To: " + contact.Email)
			continue
		}

		//Log email sent
		fmt.Println("Email Sent To: " + contact.Email)

		select {
		case <-time.After(time.Second * 3): // Wait 3 seconds before sending each email
			e.InProgress = false
			if len(*template.GetRecipients()) <= 0 { //Check if any emails are left to send to
				return nil
			}
		}
	}

}
