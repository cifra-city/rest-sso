package accessbox

import "github.com/cifra-city/rest-sso/pkg/mailman/meta"

func (a *Service) GetAndDeleteOperation(email string, operationType string, metadata meta.Data) *Data {
	a.mu.Lock()
	defer a.mu.Unlock()

	if operations, exists := a.usersList[email]; exists {
		for i, operation := range operations {
			if operation.OperationType == operationType &&
				operation.Meta.IP == metadata.IP &&
				operation.Meta.UserAgent == metadata.UserAgent {
				a.usersList[email] = append(operations[:i], operations[i+1:]...)
				if len(a.usersList[email]) == 0 {
					delete(a.usersList, email)
				}
				return &operation
			}
		}
	}

	return nil
}
