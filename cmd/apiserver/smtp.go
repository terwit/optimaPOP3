package main

import (
	"fmt"
	"net/smtp"
	"os"
)



func SendMail() {
	// user we are authorizing as
	from := "vitkhv@mail.ru"
	
	// use we are sending email to
	to := "vitkhv@gmail.com"
	
	// server we are authorized to send email through
	host := "smtp.mail.ru"
	
	// Create the authentication for the SendMail()
	// using PlainText, but other authentication methods are encouraged
	auth := smtp.PlainAuth("vitkhv", from, "tvv@2012", host)
	
	// NOTE: Using the backtick here ` works like a heredoc, which is why all the
	// rest of the lines are forced to the beginning of the line, otherwise the
	// formatting is wrong for the RFC 822 style
	message := `
	Проба пера
`
	
	if err := smtp.SendMail(host+":25", auth, from, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
		os.Exit(1)
	}
	fmt.Println("Email Sent!")
}