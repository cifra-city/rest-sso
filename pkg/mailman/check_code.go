package mailman

func (m *Mailman) CheckCode(username string, code string, operationType string) bool {
	return m.Mailbox.CheckInBox(username, code, operationType)
}
