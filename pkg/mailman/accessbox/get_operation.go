package accessbox

func (a *Service) GetSuccessForUser(email string, operationType string) *Data {
	a.mu.Lock()
	defer a.mu.Unlock()

	if operations, exists := a.usersList[email]; exists {
		for _, operation := range operations {
			if operation.OperationType == operationType {
				return &operation
			}
		}
	}

	return nil
}
