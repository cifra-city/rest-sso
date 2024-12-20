package mailman

func (m *Mailman) CheckCode(email string, operationType string, code string, userAgent string, IP string) error {
	user := m.Mailbox.CheckUserInBox(email, operationType, code)
	if user == nil {
		return ErrNotFound
	}

	if user.Meta.UserAgent != userAgent || user.Meta.IP != IP {
		return ErrAccessDenied
	}

	return nil
}
