package mailman

func (m *Mailman) CheckAndDeleteCode(email string, code string, operationType string) bool {
	return m.Mailbox.CheckAndDeleteInBox(email, code, operationType)
}
