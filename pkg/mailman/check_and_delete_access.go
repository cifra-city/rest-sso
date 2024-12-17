package mailman

func (m *Mailman) CheckAndDeleteAccessForUser(email string, operationType string) bool {
	return m.AccessBox.GetAndDeleteOperation(email, operationType)
}
