package mailbox

import "github.com/sirupsen/logrus"

func (m *Mailbox) CheckUserInBox(username string, operationType string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[username]; exists {
		if _, opExists := operations[operationType]; opExists {
			logrus.Infof("Code for user '%s' and operation '%s' is correct and has been used", username, operationType)
		} else {
			logrus.Warnf("No operation '%s' found for user '%s'", operationType, username)
		}
	} else {
		logrus.Warnf("No codes found for user '%s'", username)
	}

	return false
}
