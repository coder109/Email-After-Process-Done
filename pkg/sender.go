package pkg

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(content string, sender string, recver string, port int, smtp_server string, password string, attach string, is_failed bool) {
	e := email.NewEmail()
	e.From = sender
	e.To = []string{recver}
	if is_failed {
		e.Subject = "Error from Server Profiler"
	} else {
		e.Subject = "Success Information from Server Profiler"
	}
	e.Text = []byte(content)
	_, attach_err := e.AttachFile(attach)
	if attach_err != nil {
		panic(attach_err)
	}
	err := e.Send(smtp_server+":"+fmt.Sprint(port), smtp.PlainAuth("", sender, password, smtp_server))
	if err != nil {
		panic(err)
	}
}
