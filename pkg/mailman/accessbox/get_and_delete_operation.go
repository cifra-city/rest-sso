package accessbox

func (a *AccessBox) GetAndDeleteOperation(username, operation string) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if operations, exists := a.usersList[username]; exists {
		for i, op := range operations {
			if op == operation {
				a.usersList[username] = append(operations[:i], operations[i+1:]...)
				if len(a.usersList[username]) == 0 {
					delete(a.usersList, username)
				}
				return true
			}
		}
	}
	return false
}
