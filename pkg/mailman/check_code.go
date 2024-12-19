package mailman

import (
	"errors"
)

func (m *Mailman) CheckCode(email string, operationType string, code string, userAgent string, IP string) error {
	user, _ := m.Mailbox.CheckUserInBox(email, operationType, code)
	if user == nil {
		return errors.New("not found")
	}

	if user.Meta.UserAgent != userAgent || user.Meta.IP != IP {
		return errors.New("access denied")
	}

	return nil
}
