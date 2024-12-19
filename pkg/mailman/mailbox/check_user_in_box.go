package mailbox

import (
	"github.com/cifra-city/rest-sso/pkg/mailman/crypto"
	"github.com/sirupsen/logrus"

	"log"
)

func (m *Service) CheckUserInBox(username string, operationType string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[username]; exists {
		if encryptedData, opExists := operations[operationType]; opExists {
			_, err := crypto.Decrypt(encryptedData.ConfidenceCode, string(m.key))
			if err != nil {
				log.Printf("Failed to decrypt data for user '%s' and operation '%s': %v", username, operationType, err)
				return false
			}
			logrus.Infof("Code for user '%s' and operation '%s' exists", username, operationType)
			return true
		} else {
			logrus.Warnf("No operation '%s' found for user '%s'", operationType, username)
		}
	} else {
		logrus.Warnf("No codes found for user '%s'", username)
	}

	return false
}
