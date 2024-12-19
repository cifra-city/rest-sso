package accessbox

func (a *Service) GetOperationsForUser(email string) []Data {
	a.mu.Lock()
	defer a.mu.Unlock()

	if operations, exists := a.usersList[email]; exists {
		return operations
	}

	return nil
}
