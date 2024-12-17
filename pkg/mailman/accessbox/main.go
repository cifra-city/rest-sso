package accessbox

import "sync"

type AccessBox struct {
	mu        *sync.Mutex
	usersList map[string][]string // username -> operationType
}

func NewAccessBox() *AccessBox {
	return &AccessBox{
		mu:        &sync.Mutex{},
		usersList: make(map[string][]string),
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
