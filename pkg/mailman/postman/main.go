package postman

type Postman struct {
	Port     string
	Host     string
	Address  string
	Password string
}

func NewPostman(port string, host string, address string, password string) *Postman {
	return &Postman{
		Port:     port,
		Host:     host,
		Address:  address,
		Password: password,
	}
}
