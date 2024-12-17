package mailbox

import (
	"sync"
)

type Data struct {
	ConfidenceCode string
	OperationType  string
}

type Mailbox struct {
	mu       *sync.Mutex
	listCode map[string]map[string]Data // username -> operationType -> Data
}

func NewMailbox() *Mailbox {
	return &Mailbox{
		mu:       &sync.Mutex{},
		listCode: make(map[string]map[string]Data),
	}
}
