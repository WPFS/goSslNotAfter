package mail

import (
	"encoding/json"
	"net/smtp"
	"os"
	"strings"
)

type cfg struct {
	MailHost string
	MailPort string
	From     string
	Password string
	To       string
	Subject  string
}

func Sendmail(body string) {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	conf := cfg{}

	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	//fmt.Println(conf.MailHost, conf.MailPort, conf.From, conf.Password, conf.To, conf.Subject)
	auth := smtp.PlainAuth("", conf.From, conf.Password, conf.MailHost)

	msg := []byte("To: " + conf.To + "\r\n" +
		"Subject: " + conf.Subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	toUser := strings.Split(conf.To, ";")
	err = smtp.SendMail(conf.MailHost+":"+conf.MailPort, auth, conf.From, toUser, msg)
	if err != nil {
		panic(err)
	}
}
