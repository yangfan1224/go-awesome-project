package gomail

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail() error{
	m := gomail.NewMessage()
	m.SetHeader("From", "dmp-alarm@sndo.com")
	m.SetHeader("To", "yangfan@sndo.com", "80476178@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.exmail.qq.com", 25, "xxxx@qq.com", "ssssx")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("send mail failed. error is %s\n", err)
		return err
	}
	return nil
}
