package mailbox

import (
	"sync"
	"time"

	"github.com/cifra-city/rest-sso/pkg/mailman/meta"
)

type Data struct {
	ConfidenceCode string
	OperationType  string
	Meta           meta.Data
}

type Mailbox interface {
	AddToBox(email string, ConfidenceCode string, operationType string, metadata meta.Data, seconds time.Duration) error
	GetUserData(email string, operationType string) *Data
	CheckUserInBox(username string, operationType string) bool
	DeleteFromBox(email string, operationType string) error
	CheckAndDeleteInBox(username string, ConfidenceCode string, operationType string, UserAgent string, IP string) bool
	Drop()
}

type Service struct {
	mu       *sync.Mutex
	listCode map[string]map[string]Data // email -> operationType -> Data
	key      []byte
}

func NewMailbox(key []byte) *Service {
	return &Service{
		mu:       &sync.Mutex{},
		listCode: make(map[string]map[string]Data),
		key:      key,
	}
}
