package mail

import (
	"os/exec"
	"os"
	"io"
	"gopkg.in/gomail.v2"
)

const sendmail_executable = "/usr/sbin/postfix"

type Email struct {
	ID		   int      `jsonapi:"primary,emails"`
	Subject    string   `jsonapi:"attr,subject"`
	Body	   string   `jsonapi:"attr,body"`
	Sender     string   `jsonapi:"attr,sender"`
	Recipients []string `jsonapi:"attr,addrs"`
}

type EmailSender interface {
	Send(msg *EmailMsg) error
}

type SmtpClient struct {
	Server   string
	Port     int
	Username string
	Password string
}

type SendmailClient struct {}

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

func (cl *SmtpClient) Send(m *Email) error {
	msg := NewMessage()
	msg.From(m.Sender)
	msg.To(m.Recipients ...)
	msg.Subject(m.Subject)
	msg.Body(m.Body)
	d := gomail.NewDialer(cl.Server, cl.Port, cl.Username, cl.Password)
	if err := d.DialAndSend(&msg.Message); err != nil {
		return err
	}
	return nil
}

func (cl *SendmailClient) send(m *EmailMsg) error {
	s := gomail.SendFunc(func(from string, to []string, m io.WriterTo) error {
		return func(m io.WriterTo) error {
			cmd := exec.Command(sendmail_executable, "-t")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if pw, err := cmd.StdinPipe(); err == nil {
				defer pw.Close()
				err = cmd.Start()
				if err != nil {
					return err
				}
				_, err = m.WriteTo(pw)
				if err != nil {
					return err
				}
				err = cmd.Wait()
				if err != nil {
					return err
				}
			}
			return nil
		}(m)
	})
	return gomail.Send(s, &m.Message)
}
