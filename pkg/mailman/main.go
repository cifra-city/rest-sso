package mailman

import (
	"github.com/cifra-city/rest-sso/pkg/mailman/accessbox"
	"github.com/cifra-city/rest-sso/pkg/mailman/mailbox"
	"github.com/cifra-city/rest-sso/pkg/mailman/postman"
)

type Mailman struct {
	Mailbox   *mailbox.Service
	AccessBox *accessbox.Service
	Postman   *postman.Postman
}

func NewMailman(port string, host string, address string, password string, key []byte) *Mailman {
	return &Mailman{
		Mailbox:   mailbox.NewMailbox(key),
		Postman:   postman.NewPostman(port, host, address, password),
		AccessBox: accessbox.NewAccessBox(),
	}
}
