package mailman

func (m *Mailman) CheckCode(email string, code string, operationType string) bool {
	return m.Mailbox.CheckInBox(email, code, operationType)
}
