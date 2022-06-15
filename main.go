package main

import (
	"fmt"

	. "github.com/mailjet/mailjet-apiv3-go"
)

func main() {
	mailjetClient := NewMailjetClient("83a241161ed94c0d8668ce984a624a4d", "76cc38cf83e14ef68477ecdd643d5363")
	email := &InfoSendMail{
		FromEmail: "gozza15bdg@gmail.com",
		FromName:  "Mailjet Pilot",
		Subject:   "Your email flight plan!",
		TextPart:  "Dear passenger, welcome to Mailjet! May the delivery force be with you!",
		HTMLPart:  "<h3>Dear passenger, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!<br />May the delivery force be with you!",
		Recipients: []Recipient{
			Recipient{
				Email: "gozzafadillah@gmail.com",
			},
		},
	}
	res, err := mailjetClient.SendMail(email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}
