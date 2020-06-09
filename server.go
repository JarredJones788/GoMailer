package main

import (
	"bufio"
	"mailer"
	"os"
	"templates"
)

func main() {
	mail := mailer.Emailer{}.Init()

	jarred := mailer.Contact{"Jarred", "jarred788@hotmail.com"}
	scopes := mailer.Contact{"Scopes", "iamscopes123@hotmail.com"}

	template := templates.TestTemplate{}.Create([]mailer.Contact{jarred, scopes})
	go mail.SendEmails(template)

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
