package mailbox

import (
	"log"

	"github.com/cifra-city/rest-sso/pkg/mailman/crypto"
)

func (m *Service) GetUserData(email string, operationType string) *Data {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if encryptedData, opExists := operations[operationType]; opExists {
			// Дешифруем ConfidenceCode
			decryptedCode, err := crypto.Decrypt(encryptedData.ConfidenceCode, string(m.key))
			if err != nil {
				log.Printf("Failed to decrypt data for email '%s' and operation '%s': %v", email, operationType, err)
				return nil
			}
			// Возвращаем дешифрованные данные
			return &Data{
				ConfidenceCode: decryptedCode,
				OperationType:  encryptedData.OperationType,
				Meta:           encryptedData.Meta,
			}
		}
	}

	return nil
}
