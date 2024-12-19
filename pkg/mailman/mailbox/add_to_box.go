package mailbox

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cifra-city/rest-sso/pkg/mailman/crypto"
	"github.com/cifra-city/rest-sso/pkg/mailman/meta"
)

func (m *Service) AddToBox(email string, ConfidenceCode string, operationType string, metadata meta.Data, seconds time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.listCode[email]; !exists {
		m.listCode[email] = make(map[string]Data)
	}

	if _, exists := m.listCode[email][operationType]; exists {
		return errors.New("code already exists")
	}

	// Создаем структуру данных
	data := Data{
		ConfidenceCode: ConfidenceCode,
		OperationType:  operationType,
		Meta:           metadata,
	}

	// Преобразуем структуру в JSON
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Шифруем JSON данные
	encryptedData, err := crypto.Encrypt(string(dataJSON), string(m.key))
	if err != nil {
		return err
	}

	// Сохраняем зашифрованные данные
	m.listCode[email][operationType] = Data{
		ConfidenceCode: encryptedData,
		OperationType:  operationType,
		Meta:           metadata, // Метаданные сохраняются без шифрования (если требуется, это можно изменить)
	}

	// Удаляем данные через заданное время
	time.AfterFunc(seconds*time.Second, func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		if operations, exists := m.listCode[email]; exists {
			if _, opExists := operations[operationType]; opExists {
				delete(operations, operationType)

				if len(operations) == 0 {
					delete(m.listCode, email)
				}
			}
		}
	})

	return nil
}
