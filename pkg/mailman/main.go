package mailman

import (
	"github.com/cifra-city/rest-sso/pkg/mailman/accessbox"
	"github.com/cifra-city/rest-sso/pkg/mailman/mailbox"
	"github.com/cifra-city/rest-sso/pkg/mailman/postman"
)

type Mailman struct {
	Mailbox   *mailbox.Mailbox
	AccessBox *accessbox.AccessBox
	Postman   *postman.Postman
}

func NewMailman(port string, host string, address string, password string) *Mailman {
	return &Mailman{
		Mailbox:   mailbox.NewMailbox(),
		Postman:   postman.NewPostman(port, host, address, password),
		AccessBox: accessbox.NewAccessBox(),
	}
}
