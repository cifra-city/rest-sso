package mailman

import (
	"errors"
)

func (m *Mailman) CheckAccess(email string, operationType string, userAgent string, IP string) error {
	access := m.AccessBox.GetSuccessForUser(email, operationType)
	if access == nil {
		return errors.New("not found")
	}
	if access.Meta.IP == IP && access.Meta.UserAgent == userAgent {
		return nil // Доступ предоставлен
	}
	return errors.New("access denied")
}
