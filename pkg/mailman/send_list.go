package mailman

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (m *Mailman) SendList(email string, operationType string, templateList string, seconds time.Duration) error {
	code, err := m.Postman.GenCode()
	if err != nil {
		logrus.Errorf("Failed to generate code: %v", err)
		return err
	}
	m.Mailbox.AddToBox(email, code, operationType, seconds)
	err = m.Postman.SendCode(email, code, templateList)
	if err != nil {
		logrus.Errorf("Failed to send code: %v", err)
		return err
	}

	return nil
}
