package accessbox

import "time"

func (a *AccessBox) AddAccess(email string, operation string, minutes time.Duration) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !contains(a.usersList[email], operation) {
		a.usersList[email] = append(a.usersList[email], operation)
	}

	go func() {
		<-time.After(minutes * time.Minute)
		a.DeleteOperationForUser(email, operation)
	}()
}
