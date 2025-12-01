package pkg

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(content string) {
	e := email.NewEmail()
	e.From = "<SENDER>"
	e.To = []string{"<RECVER>"}
	e.Subject = "<SUBJECT>"
	e.Text = []byte(content)
	err := e.Send("<SENDER SERVER>:<PORT>", smtp.PlainAuth("", "<SENDER>", "<VERIFICATION CODE>", "<SENDER SERVER>"))
	if err != nil {
		panic(err)
	}
}
