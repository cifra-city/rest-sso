package accessbox

import (
	"log"
	"time"

	"github.com/cifra-city/rest-sso/pkg/mailman/meta"
)

func (a *Service) AddAccessForOperation(email string, operationType string, metadata meta.Data, minutes time.Duration) {
	a.mu.Lock()
	defer a.mu.Unlock()

	for _, operation := range a.usersList[email] {
		if operation.OperationType == operationType {
			log.Printf("Operation '%s' for user '%s' already exists", operationType, email)
			return
		}
	}

	a.usersList[email] = append(a.usersList[email], Data{
		OperationType: operationType,
		Meta:          metadata,
	})

	go func() {
		<-time.After(minutes * time.Minute)
		a.DeleteOperationForUser(email, operationType)
	}()
}
