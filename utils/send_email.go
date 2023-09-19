package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func Send(body string, to string) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	email := os.Getenv("Email")
	emailPwd := os.Getenv("Email_PWD")
	from := email
	pass := emailPwd

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Email Verification\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Successfully sended to " + to)
}
