package accessbox

import "time"

func (a *AccessBox) AddOperation(email string, operation string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !contains(a.usersList[email], operation) {
		a.usersList[email] = append(a.usersList[email], operation)
	}

	go func() {
		<-time.After(15 * time.Minute)
		a.DeleteOperationForUser(email, operation)
	}()
}
