package postman

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func (p *Postman) SendCode(to string, code string, templateList string) error {
	mes := gomail.NewMessage()
	mes.SetHeader("From", p.Address)
	mes.SetHeader("To", to)
	mes.SetHeader("Subject", "Email Confirmation")

	templatePath := filepath.Join(templateList)
	tmplContent, err := ioutil.ReadFile(templatePath)
	if err != nil {
		logrus.Errorf("Error reading email template: %v", err)
		return err
	}

	tmpl, err := template.New("emailTemplate").Parse(string(tmplContent))
	if err != nil {
		logrus.Errorf("Error parsing email template: %v", err)
		return err
	}

	var bodyContent bytes.Buffer
	err = tmpl.Execute(&bodyContent, map[string]string{"Code": code})
	if err != nil {
		logrus.Errorf("Error executing template: %v", err)
		return err
	}

	mes.SetBody("text/html", bodyContent.String())

	smtpPort, err := strconv.Atoi(p.Port)
	if err != nil {
		logrus.Fatalf("Invalid SMTP port: %v", err)
	}

	d := gomail.NewDialer(p.Host, smtpPort, p.Address, p.Password)

	if err := d.DialAndSend(mes); err != nil {
		return err
	}

	logrus.Debugf("Code sent to %s", to)
	return nil
}
