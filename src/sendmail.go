package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type (
	// Payload ...
	Payload struct {
		// application/system name
		Sender    string
		Title     string
		Message   string
		IsHTML    bool
		Template  string
		UserIP    string
		UserAgent string
	}
)

var (
	sa   smtp.Auth
	host string
	port string
	from string
	pass string
	to   string
)

func init() {
	host = os.Getenv("GOMAIL_SRV_HOST")
	port = os.Getenv("GOMAIL_SRV_PORT")
	from = os.Getenv("GOMAIL_AUTH_USERNAME")
	pass = os.Getenv("GOMAIL_AUTH_PASSWORD")
	to = os.Getenv("GOMAIL_TO")
	sa = smtp.PlainAuth("", from, pass, host)
	if os.Getenv("DEBUG") != "" {
		log.Println(host, port, from, pass, to, sa)
	}
}

// SendMail ...
func SendMail(p Payload) error {
	msg := &bytes.Buffer{}
	msg.WriteString("From: " + p.Sender + " <" + from + ">\n")
	msg.WriteString("To: " + to + "\n")
	msg.WriteString("Subject: " + p.Title + "\n")
	msg.WriteString(fmt.Sprintf("MIME-version: 1.0;\nContent-Type: %s; charset=\"UTF-8\";\n\n",
		Ternary(p.IsHTML, "text/html", "text/plain")))
	msg.WriteString(getTemplate(p))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", host, port), sa, from, strings.Split(to, ";"), msg.Bytes())
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}
	log.Printf("Sent! %s | %#v\n", to, p)
	return nil
}

func getTemplate(p Payload) string {
	if p.Template == "" {
		p.Template = "default"
	}

	s := fmt.Sprintf("templates/%s.%s", p.Template, Ternary(p.IsHTML, "html", "txt"))

	f, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("tpl").Parse(string(f))
	check(err)

	var b bytes.Buffer
	err = t.Execute(&b, p)
	check(err)

	return b.String()
}
