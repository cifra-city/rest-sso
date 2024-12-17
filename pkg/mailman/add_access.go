package mailman

func (m *Mailman) AddAccessForUser(email string, operationType string) {
	m.AccessBox.AddOperation(email, operationType)
}
