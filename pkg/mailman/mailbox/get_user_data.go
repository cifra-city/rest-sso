package mailbox

func (m *Service) GetUserData(email string, operationType string) *Data {
	m.mu.Lock()
	defer m.mu.Unlock()

	if operations, exists := m.listCode[email]; exists {
		if data, opExists := operations[operationType]; opExists {
			return &Data{
				ConfidenceCode: data.ConfidenceCode,
				OperationType:  data.OperationType,
				Meta:           data.Meta,
			}
		}
	}

	return nil
}
