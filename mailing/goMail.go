package mailing

import (
	"log"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"
)

//SendEmail Send Email
func SendEmail(recipients []string, subject string, htmlBody string) {

	configEmail := os.Getenv("MAIL_USER")
	configPassword := os.Getenv("MAIL_PASS")
	configHost := os.Getenv("MAIL_HOST")
	configPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", configEmail)

	addresses := make([]string, len(recipients))
	for i, recipient := range recipients {
		addresses[i] = m.FormatAddress(recipient, "")
	}

	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewPlainDialer(configHost, configPort, configEmail, configPassword)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	log.Println("Mail sent!")
}
