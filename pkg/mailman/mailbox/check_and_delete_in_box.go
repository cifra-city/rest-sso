package mailbox

import (
	"github.com/cifra-city/rest-sso/pkg/mailman/crypto"
	"github.com/sirupsen/logrus"
)

func (m *Service) CheckAndDeleteInBox(username string, ConfidenceCode string, operationType string, UserAgent string, IP string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Проверяем наличие пользователя в мапе
	if operations, exists := m.listCode[username]; exists {
		// Проверяем наличие операции для пользователя
		if data, opExists := operations[operationType]; opExists {
			// Расшифровываем код
			decryptedCode, err := crypto.Decrypt(data.ConfidenceCode, string(m.key))
			if err != nil {
				logrus.Errorf("Failed to decrypt ConfidenceCode for user '%s' and operation '%s': %v", username, operationType, err)
				return false
			}

			// Проверяем соответствие данных
			if decryptedCode == ConfidenceCode && data.Meta.UserAgent == UserAgent && data.Meta.IP == IP {
				// Удаляем операцию
				delete(operations, operationType)
				logrus.Infof("Code for user '%s' and operation '%s' is correct and has been used", username, operationType)

				// Удаляем пользователя, если у него больше нет операций
				if len(operations) == 0 {
					delete(m.listCode, username)
				}

				return true
			}

			logrus.Warnf("Incorrect data for user '%s' and operation '%s'. Expected: Code='%s', UserAgent='%s', IP='%s'. Got: Code='%s', UserAgent='%s', IP='%s'",
				username, operationType, decryptedCode, data.Meta.UserAgent, data.Meta.IP, ConfidenceCode, UserAgent, IP)
		} else {
			logrus.Warnf("No operation '%s' found for user '%s'", operationType, username)
		}
	} else {
		logrus.Warnf("No codes found for user '%s'", username)
	}

	return false
}
