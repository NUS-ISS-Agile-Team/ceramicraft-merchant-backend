package email

import (
	"sync"
	"gopkg.in/mail.v2"

	"github.com/cerami-craft-shop/merchant-backend/config"
)

type EmailSender struct {
	SmtpHost      string `json:"smtp_host"`
	SmtpEmailFrom string `json:"smtp_email_from"`
	SmtpPass      string `json:"smtp_pass"`
}

var (
	instance *EmailSender
	once     sync.Once
)

// GetInstance 获取EmailSender单例实例
func GetInstance() *EmailSender {
	once.Do(func() {
		eConfig := config.GetConfig().Email
		instance = &EmailSender{
			SmtpHost:      eConfig.SmtpHost,
			SmtpEmailFrom: eConfig.SmtpEmail,
			SmtpPass:      eConfig.SmtpPass,
		}
	})
	return instance
}

// NewEmailSender 保留原有函数以兼容现有代码，但内部使用单例模式
func NewEmailSender() *EmailSender {
	return GetInstance()
}

// Send 发送邮件
func (s *EmailSender) Send(data, emailTo, subject string) error {
	m := mail.NewMessage()
	m.SetHeader("From", s.SmtpEmailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", data)
	d := mail.NewDialer(s.SmtpHost, 465, s.SmtpEmailFrom, s.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
