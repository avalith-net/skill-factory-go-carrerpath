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
	from := mail.Address{"CareerPath", "careerpath.avalith@gmail.com"}
	to := mail.Address{userName, email} //cambiar dependiendo del usuario que quiera recuperar la contraseña
	subject := "Recuperacion de Contraseña"

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

	t, err := template.ParseFiles("templateMail.html")
	checkError(err, "")

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkError(err, "") // agregar errores

	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "careerpath.avalith@gmail.com", "careerpath1", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkError(err, "")

	client, err := smtp.NewClient(conn, host)
	checkError(err, "")

	err = client.Auth(auth)
	checkError(err, "")

	err = client.Mail(from.Address)
	checkError(err, "")

	err = client.Rcpt(to.Address)
	checkError(err, "")

	w, err := client.Data()
	checkError(err, "")

	_, err = w.Write([]byte(message))
	checkError(err, "")

	err = w.Close()
	checkError(err, "")

	client.Quit()

	return "Email enviado con exito.", nil

}
