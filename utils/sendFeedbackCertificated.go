package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/alecthomas/template"
	"net/mail"
	"net/smtp"
)

type Dest struct {
	Name        string
	Feedback string
}

func SendFeedbackCertificated(userName string, email string, Feedback string) error {
	from := mail.Address{Name: "CareerPath", Address: "careerpath.avalith@gmail.com"}
	to := mail.Address{Name: userName, Address: email}
	subject := "Feedback of certification Skill"

	Dest := Dest{Name: to.Address, Feedback: Feedback }

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""

	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("utils/assets/feedbackTemplate.html")
	CheckError(err, "error in the template html")

	buf := new(bytes.Buffer)
	err = t.Execute(buf, Dest)
	CheckError(err, "error in the execute")

	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "careerpath.avalith@gmail.com", "careerpath1", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	CheckError(err, "error in the tlsConfig")

	client, err := smtp.NewClient(conn, host)
	CheckError(err, "error in the client create")

	err = client.Auth(auth)
	CheckError(err, "error in the authorization")

	err = client.Mail(from.Address)
	CheckError(err, "error in the address (from)")

	err = client.Rcpt(to.Address)
	CheckError(err, "error in the address (to)")

	w, err := client.Data()

	_, err = w.Write([]byte(message))
	CheckError(err, "error in the write")

	err = w.Close()
	CheckError(err, "error in the close")

	client.Quit()

	return  nil

}