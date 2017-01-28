package main

import (
	"os/exec"
	"os"
	"io"
	"gopkg.in/gomail.v2"
)

const executable = "/usr/sbin/postfix"

type EmailMsg struct {
	gomail.Message
}

func NewMessage() EmailMsg {
	return EmailMsg{
		Message: *gomail.NewMessage(),
	}
}

func (msg *EmailMsg) To(recipient ... string) {
	msg.Message.SetHeader("To", recipient ...)
}

func (msg *EmailMsg) From(sender string) {
	msg.SetHeader("From", sender)
}

func (msg *EmailMsg) Subject(subject string) {
	msg.SetHeader("Subject", subject)
}

func (msg *EmailMsg) Body(body string) {
	const TextPlain = "text/plain"
	msg.SetBody(TextPlain, body)
}

type SMTPConfig struct {
	Server   string  //"sendmail"
	Port     int     // 25
	Username string  // "user"
	Password string  // "123456"
}

func (msg *EmailMsg) SendSmtp(conf *SMTPConfig) error {
	d := gomail.NewDialer(conf.Server, conf.Port, conf.Username, conf.Password)
	if err := d.DialAndSend(&msg.Message); err != nil {
		return err
	}
	return nil
}

func (msg *EmailMsg) SendMail() error {
	s := gomail.SendFunc(func(from string, to []string, m io.WriterTo) error {
		return sendmail(m)
	})
	return gomail.Send(s, &msg.Message)
}

func sendmail(m io.WriterTo) (err error) {
	cmd := exec.Command(executable, "-t")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	pw, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	_, err = m.WriteTo(pw)
	if err != nil {
		return
	}
	err = pw.Close()
	if err != nil {
		return
	}
	err = cmd.Wait()
	if err != nil {
		return
	}
	return nil
}
