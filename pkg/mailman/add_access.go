package mailman

import (
	"time"

	"github.com/cifra-city/rest-sso/pkg/mailman/meta"
)

func (m *Mailman) AddAccess(email string, operationType string, userAgent string, IP string, minutes time.Duration) {
	m.AccessBox.AddAccessForOperation(email, operationType, meta.Data{
		UserAgent: userAgent,
		IP:        IP,
	}, minutes)
}
