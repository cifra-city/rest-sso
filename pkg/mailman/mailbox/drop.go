package mailbox

import "github.com/sirupsen/logrus"

func (m *Service) Drop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.listCode = make(map[string]map[string]Data)
	logrus.Info("Service has been cleared")
}
