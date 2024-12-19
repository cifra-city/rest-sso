package mailbox

import (
	"github.com/sirupsen/logrus"
)

func (m *Service) CheckUserInBox(email string, operationType string, ConfidenceCode string) (*Data, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if data, opExists := operations[operationType]; opExists {
			if data.ConfidenceCode == ConfidenceCode {
				logrus.Infof("Code for user '%s' and operation '%s' is correct and has been used", email, operationType)
				return &data, nil
			} else {
				logrus.Warnf("Code for user '%s' and operation '%s' is incorrect", email, operationType)
			}
		} else {
			logrus.Warnf("No operation '%s' found for user '%s'", operationType, email)
		}
	} else {
		logrus.Warnf("No codes found for user '%s'", email)
	}

	return nil, nil
}
