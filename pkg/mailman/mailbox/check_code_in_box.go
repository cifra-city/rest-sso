package mailbox

import "github.com/sirupsen/logrus"

func (m *Mailbox) CheckInBox(email string, ConfidenceCode string, operationType string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if data, opExists := operations[operationType]; opExists {
			if ConfidenceCode == data.ConfidenceCode {
				logrus.Infof("Code for user '%s' and operation '%s' is correct and has been used", email, operationType)
				return true
			}
			logrus.Warnf("Incorrect code for user '%s' and operation '%s'. Expected: '%s', Got: '%s'",
				email, operationType, data.ConfidenceCode, ConfidenceCode)
		} else {
			logrus.Warnf("No operation '%s' found for user '%s'", operationType, email)
		}
	} else {
		logrus.Warnf("No codes found for user '%s'", email)
	}

	return false
}
