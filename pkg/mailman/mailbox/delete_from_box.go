package mailbox

import "errors"

func (m *Service) DeleteFromBox(email string, operationType string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if _, opExists := operations[operationType]; opExists {
			delete(operations, operationType)

			if len(operations) == 0 {
				delete(m.listCode, email)
			}

			return nil
		}
	}

	return errors.New("no data found")
}
