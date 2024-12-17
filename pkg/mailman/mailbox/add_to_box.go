package mailbox

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (m *Mailbox) AddToBox(username string, ConfidenceCode string, operationType string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.listCode[username]; !exists {
		m.listCode[username] = make(map[string]Data)
	}

	m.listCode[username][operationType] = Data{
		ConfidenceCode: ConfidenceCode,
		OperationType:  operationType,
	}

	time.AfterFunc(180*time.Second, func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		if operations, exists := m.listCode[username]; exists {
			if _, opExists := operations[operationType]; opExists {
				delete(operations, operationType)
				logrus.Infof("Code for user '%s' and operation '%s' has expired", username, operationType)

				if len(operations) == 0 {
					delete(m.listCode, username)
				}
			}
		}
	})
}
