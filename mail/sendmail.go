package mail

import (
	"net/smtp"
	"strings"
)

const (
	Addr     = "smtp.exmail.qq.com"
	AddrPort = "smtp.exmail.qq.com:25"
	User     = "user@qq.com"
	Password = "pwd"
	Subject  = "SSL证书马上到期, 请及时更新!!!"
)

var ToUser = "recipient@qq.com;recipient2@qq.com"

func Sendmail(body string) {
	auth := smtp.PlainAuth("", User, Password, Addr)

	msg := []byte("To: " + ToUser + "\r\n" +
		"Subject: " + Subject + "\r\n" +
		"\r\n" +
		body + "\r\n")
	toUser := strings.Split(ToUser, ";")
	err := smtp.SendMail(AddrPort, auth, User, toUser, msg)
	if err != nil {
		panic(err)
	}
}
