package mail

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"net"
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

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func validateLine(line string) error {
	if strings.ContainsAny(line, "\n\r") {
		return errors.New("smtp: A line must not contain CR or LF")
	}
	return nil
}

func SendMailSSL(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	if err := validateLine(from); err != nil {
		return err
	}
	for _, recp := range to {
		if err := validateLine(recp); err != nil {
			return err
		}
	}
	c, err := Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	if a != nil {
		if ok, _ := c.Extension("AUTH"); !ok {
			return errors.New("smtp: server doesn't support AUTH")
		}
		if err = c.Auth(a); err != nil {
			return err
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
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
	err = SendMailSSL(conf.MailHost+":"+conf.MailPort, auth, conf.From, toUser, msg)
	if err != nil {
		panic(err)
	}
}
