package accessbox

func (a *AccessBox) DeleteOperationForUser(username, operation string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if operations, exists := a.usersList[username]; exists {
		for i, op := range operations {
			if op == operation {
				// Удаляем операцию из списка
				a.usersList[username] = append(operations[:i], operations[i+1:]...)
				// Если список пуст, удаляем пользователя
				if len(a.usersList[username]) == 0 {
					delete(a.usersList, username)
				}
				return
			}
		}
	}
}
