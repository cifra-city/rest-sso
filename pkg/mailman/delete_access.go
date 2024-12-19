package mailman

func (m *Mailman) DeleteAccess(email string, operationType string) error {
	return m.AccessBox.DeleteOperationForUser(email, operationType)
}
