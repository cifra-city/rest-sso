package mailman

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (m *Mailman) SendList(email string, operationType string, templateList string, userAgent string, IP string, seconds time.Duration) error {
	code, err := m.AddCode(email, operationType, userAgent, IP, seconds)
	if err != nil {
		logrus.Errorf("Failed to add code: %v", err)
		return err
	}
	err = m.Postman.SendCode(email, *code, templateList)
	if err != nil {
		logrus.Errorf("Failed to send code: %v", err)
		return err
	}

	logrus.Debugf("Code sent to %s", email)
	return nil
}
