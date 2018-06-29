package zl_lib

import (
	"fmt"

	l4g "zl_lib/code.google.com/p/log4go"
	"zl_lib/github.com/go-gomail/gomail"
)

type Email struct {
	Cfg       Email_config
	l4gLogger *l4g.Logger
}

type Email_config struct {
	User     string
	Password string
	Host     string
	Port     int
	To       []string
}

func NewEmail(cfg Email_config, logger *l4g.Logger) *Email {
	return &Email{Cfg: cfg, l4gLogger: logger}
}

func (e *Email) SendToMail(subject, body, mailtype string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", e.Cfg.User, e.Cfg.User) // 发件人

	Headers := make(map[string][]string)
	Headers["To"] = e.Cfg.To
	m.SetHeaders(Headers)
	m.SetHeader("Subject", subject)                                                // 主题
	m.SetBody("text/html", body)                                                   // 正文
	d := gomail.NewPlainDialer(e.Cfg.Host, e.Cfg.Port, e.Cfg.User, e.Cfg.Password) // 发送邮件服务器、端口、发件人账号、发件人密码
	err := d.DialAndSend(m)
	return err
}

func (e *Email) Default_send_temp(body string) bool {
	subject := "hqs log opertion 工具报警"

	body_tmp := `
		<html>
		<body>
		<h3>
		"%s"
		</h3>
		</body>
		</html>
		`
	body = fmt.Sprintf(body_tmp, body)

	n := 0
	for {

		err := e.SendToMail(subject, body, "html")
		if err == nil {
			e.l4gLogger.Error("Send mail success!")
			return true
		}
		if n > 1 {
			e.l4gLogger.Error("Send mail error! ", err)
			return false
		}
		n++
	}
	return false
}
