package postman

import (
	"crypto/rand"
	"fmt"
)

func (p *Postman) GenCode() (string, error) {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	code := fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2]))
	return code, nil
}
