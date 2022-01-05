package gomail

import (
	"bytes"
	"html/template"
	"path/filepath"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

/***
* Author madasatya
* Send mail with gomail
*/
type MailInterface interface{
	Send() error
}

type Config struct {
	SMTP_HOST 		string 	`example:"smtp.gmail.com"`
	SMTP_PORT 		int		`example:"587"`
	SENDER_NAME 	string 	`example:"PT Admin <admin@gmail.com>"`
	AUTH_EMAIL 		string	`example:"admin@gmail.com"`
	AUTH_PASSWORD 	string	`example:"123456"`
	Params			Params
}

type Params struct{
	Subject 	string 
	To			string // email tujuan
	EmailCC 	[]string 
	TitleCC 	string 
	ViewPath	string 
	Data 		map[string]interface{}
	FileAttachment 	[]string 
}

func New(config Config) MailInterface {
	return &config
}

func (config *Config) Send() error {
	
	params := config.Params

	dir := filepath.Join(params.ViewPath)
	t := template.New(dir)

	t, err := t.ParseFiles(dir)
	if err != nil {
		return err 
	}
	
	var tpl bytes.Buffer
	if err := t.ExecuteTemplate(&tpl, filepath.Base(dir), params.Data); err != nil {
		return err 
	}

	tmplHTML := tpl.String()

    mailer := gomail.NewMessage()
	mailer.SetHeader("From", config.SENDER_NAME)
	mailer.SetHeader("To", params.To)
	
	if len(params.EmailCC) > 0 {
		mailer.SetAddressHeader("Cc", strings.Join(params.EmailCC, ","), params.TitleCC)
	}
	
	mailer.SetHeader("Subject", params.Subject)
	mailer.SetBody("text/html", tmplHTML)
	
	if params.FileAttachment != nil {
		mailer.Attach(strings.Join(params.FileAttachment, ","))
	}
	
	dialer := gomail.NewDialer(config.SMTP_HOST, config.SMTP_PORT, config.AUTH_EMAIL, config.AUTH_PASSWORD)
	
	//kirim email dengan smtp relat atau no auth/tanpa otentikasi
	//dialer := &DialAndSend(mailer)
	
	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	
	return nil
}
