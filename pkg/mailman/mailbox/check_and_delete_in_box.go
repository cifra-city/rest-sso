package mailbox

import "github.com/sirupsen/logrus"

func (m *Mailbox) CheckAndDeleteInBox(username string, ConfidenceCode string, operationType string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[username]; exists {
		if data, opExists := operations[operationType]; opExists {
			if ConfidenceCode == data.ConfidenceCode {
				delete(operations, operationType)
				logrus.Infof("Code for user '%s' and operation '%s' is correct and has been used", username, operationType)

				if len(operations) == 0 {
					delete(m.listCode, username)
				}

				return true
			}
			logrus.Warnf("Incorrect code for user '%s' and operation '%s'. Expected: '%s', Got: '%s'",
				username, operationType, data.ConfidenceCode, ConfidenceCode)
		} else {
			logrus.Warnf("No operation '%s' found for user '%s'", operationType, username)
		}
	} else {
		logrus.Warnf("No codes found for user '%s'", username)
	}

	return false
}
