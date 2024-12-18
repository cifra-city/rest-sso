package mailbox

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (m *Mailbox) AddToBox(email string, ConfidenceCode string, operationType string, seconds time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.listCode[email]; !exists {
		m.listCode[email] = make(map[string]Data)
	}

	if _, exists := m.listCode[email][operationType]; exists {
		logrus.Infof("Code for user '%s' and operation '%s' already exists", email, operationType)
		return errors.New("code already exists")
	}

	m.listCode[email][operationType] = Data{
		ConfidenceCode: ConfidenceCode,
		OperationType:  operationType,
	}

	time.AfterFunc(seconds*time.Second, func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		if operations, exists := m.listCode[email]; exists {
			if _, opExists := operations[operationType]; opExists {
				delete(operations, operationType)
				logrus.Infof("Code for user '%s' and operation '%s' has expired", email, operationType)

				if len(operations) == 0 {
					delete(m.listCode, email)
				}
			}
		}
	})

	return nil
}
