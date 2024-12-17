package mailman

import "time"

func (m *Mailman) AddAccessForUser(email string, operationType string, minutes time.Duration) {
	m.AccessBox.AddAccess(email, operationType, minutes)
}
