package mailbox

import (
	"time"

	"github.com/cifra-city/rest-sso/pkg/mailman/meta"
)

func (m *Service) AddToBox(email string, ConfidenceCode string, operationType string, metadata meta.Data, seconds time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.listCode[email]; !exists {
		m.listCode[email] = make(map[string]Data)
	}

	if _, exists := m.listCode[email][operationType]; exists {
		delete(m.listCode[email], operationType)
	}

	m.listCode[email][operationType] = Data{
		ConfidenceCode: ConfidenceCode,
		OperationType:  operationType,
		Meta:           metadata,
	}

	go func() {
		<-time.After(seconds * time.Second)
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
	}()

	return nil
}
