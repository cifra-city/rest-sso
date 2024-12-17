package mailman

func (m *Mailman) CheckAccessForUser(email string, operationType string) bool {
	accesses := m.AccessBox.GetOperationsForUser(email)
	if accesses == nil {
		return false
	}
	for _, access := range accesses {
		if access == operationType {
			return true
		}
	}
	return false
}
