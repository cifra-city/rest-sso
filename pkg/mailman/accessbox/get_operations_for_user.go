package accessbox

func (a *AccessBox) GetOperationsForUser(username string) []string {
	a.mu.Lock()
	defer a.mu.Unlock()

	if operations, exists := a.usersList[username]; exists {
		return operations
	}
	return nil
}
