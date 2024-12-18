package postman

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/sirupsen/logrus"
)

func (p *Postman) GenCode() (string, error) {
	newInt := big.NewInt(900000) // 999999 - 100000 + 1
	n, err := rand.Int(rand.Reader, newInt)
	if err != nil {
		return "", err
	}

	code := 100000 + n.Int64()

	logrus.Debugf("Generated code: %d", code)
	return fmt.Sprintf("%06d", code), nil
}
