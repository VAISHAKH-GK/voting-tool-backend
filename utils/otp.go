package utils

import (
	"fmt"
	"math/rand"
	"os"

	"gopkg.in/gomail.v2"
)

func GenerateOTP() int {
	return rand.Intn(900000) + 100000
}

func SendMail(mail string, content string) {
	var sendMailAddress = os.Getenv("MAIL_ADDRESS")
	var sendMailPassword = os.Getenv("MAIL_PASSWORD")

	fmt.Println(sendMailAddress)
	fmt.Println(sendMailPassword)

	var msg = gomail.NewMessage()
	msg.SetHeader("From", "fulgurcode@gmail.com")
	msg.SetHeader("To", mail)

	msg.SetHeader("Subject", "OTP for signup for voting tool")
	msg.SetBody("text/html", content)

	var dialer = gomail.NewDialer("smtp.gmail.com", 465, sendMailAddress, sendMailPassword)
	var err = dialer.DialAndSend(msg)
	if err != nil {
		fmt.Println(err)
	}
}
