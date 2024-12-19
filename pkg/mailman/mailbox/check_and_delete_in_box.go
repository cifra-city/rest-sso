package mailbox

import (
	"github.com/sirupsen/logrus"
)

func (m *Service) CheckAndDeleteInBox(email string, ConfidenceCode string, operationType string, UserAgent string, IP string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if data, opExists := operations[operationType]; opExists {

			if data.ConfidenceCode == ConfidenceCode && data.Meta.UserAgent == UserAgent && data.Meta.IP == IP {

				delete(operations, operationType)
				logrus.Infof("Code for user '%s' and operation '%s' is correct and has been used", email, operationType)

				if len(operations) == 0 {
					delete(m.listCode, email)
				}

				return true
			}

			logrus.Warnf("Incorrect data for user '%s' and operation '%s'. Expected: Code='%s', UserAgent='%s', IP='%s'. Got: Code='%s', UserAgent='%s', IP='%s'",
				email, operationType, ConfidenceCode, data.Meta.UserAgent, data.Meta.IP, ConfidenceCode, UserAgent, IP)
		} else {
			logrus.Warnf("No operation '%s' found for user '%s'", operationType, email)
		}
	} else {
		logrus.Warnf("No codes found for user '%s'", email)
	}

	return false
}
