package accessbox

func (a *Service) DeleteOperationForUser(email string, operationType string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if operations, exists := a.usersList[email]; exists {
		for i, operation := range operations {
			if operation.OperationType == operationType {
				a.usersList[email] = append(operations[:i], operations[i+1:]...)
				if len(a.usersList[email]) == 0 {
					delete(a.usersList, email)
				}
				return
			}
		}
	}
}
