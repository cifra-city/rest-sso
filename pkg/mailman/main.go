package mailman

import (
	"time"

	"github.com/cifra-city/rest-sso/pkg/mailman/accessbox"
	"github.com/cifra-city/rest-sso/pkg/mailman/mailbox"
	"github.com/cifra-city/rest-sso/pkg/mailman/postman"
)

type Mailman struct {
	Mailbox   *mailbox.Service
	AccessBox *accessbox.Service
	Postman   *postman.Postman
}

func NewMailman(port string, host string, address string, password string) *Mailman {
	return &Mailman{
		Mailbox:   mailbox.NewMailbox(),
		Postman:   postman.NewPostman(port, host, address, password),
		AccessBox: accessbox.NewAccessBox(),
	}
}

type MailmanInterface interface {
	AddAccess(email string, operationType string, userAgent string, IP string, minutes time.Duration)
	AddCode(email string, operationType string, userAgent string, IP string, seconds time.Duration) (*string, error)

	CheckAccess(email string, operationType string, userAgent string, IP string) error
	CheckCode(email string, operationType string, code string, userAgent string, IP string) error

	DeleteCode(email string, operationType string) error
	DeleteAccess(email string, operationType string) error

	SendList(email string, operationType string, templateList string, userAgent string, IP string, seconds time.Duration) error
}
