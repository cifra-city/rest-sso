package accessbox

import (
	"sync"
	"time"

	"github.com/cifra-city/rest-sso/pkg/mailman/meta"
)

type Data struct {
	OperationType string
	Meta          meta.Data
}

type Service struct {
	mu        *sync.Mutex
	usersList map[string][]Data // username -> operationType
}

type AccessBox interface {
	AddAccessForOperation(email string, operationType string, metadata meta.Data, minutes time.Duration)
	DeleteOperationForUser(email string, operation string)
	GetAndDeleteOperation(email string, operation string, metadata meta.Data) *Data
	GetOperation(email string, operationType string) *Data
	GetOperationsForUser(email string) []Data
}

func NewAccessBox() *Service {
	return &Service{
		mu:        &sync.Mutex{},
		usersList: make(map[string][]Data),
	}
}

func contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
