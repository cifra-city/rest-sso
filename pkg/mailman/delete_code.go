package mailman

func (m *Mailman) DeleteCode(email string, operationType string) error {
	return m.Mailbox.DeleteFromBox(email, operationType)
}
