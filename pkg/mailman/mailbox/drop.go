package mailbox

import "github.com/sirupsen/logrus"

func (m *Mailbox) Drop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.listCode = make(map[string]map[string]Data)
	logrus.Info("Mailbox has been cleared")
}
