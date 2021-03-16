package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
)

type Dest struct {
	Name string
}

func checkError(err error, message string) {
	if err != nil {
		AppendError(message)
	}
}

func PasswordRecovery(userName string, email string) (string, error) {
	from := mail.Address{Name: "CareerPath", Address: "careerpath.avalith@gmail.com"}
	to := mail.Address{Name: userName, Address: email}
	subject := "Password Recovery"

	dest := Dest{Name: to.Address}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""

	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("utils/assets/templateMail.html")
	checkError(err, "error in the template html")

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkError(err, "error in the execute")

	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "careerpath.avalith@gmail.com", "careerpath1", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkError(err, "error in the tlsConfig")

	client, err := smtp.NewClient(conn, host)
	checkError(err, "error in the client create")

	err = client.Auth(auth)
	checkError(err, "error in the authorization")

	err = client.Mail(from.Address)
	checkError(err, "error in the address (from)")

	err = client.Rcpt(to.Address)
	checkError(err, "error in the address (to)")

	w, err := client.Data()

	_, err = w.Write([]byte(message))
	checkError(err, "error in the write")

	err = w.Close()
	checkError(err, "error in the close")

	client.Quit()

	return "Email successfully sent,  please check your inbox for a link to complete the reset.", nil

}
