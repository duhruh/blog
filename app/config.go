package app

type Config interface {
	GetHTTPBindAddress() string
}
type config struct {
	httpBindAddress string
}

func NewConfig(address string) Config {
	return config{httpBindAddress: address}
}

func (c config) GetHTTPBindAddress() string {
	return c.httpBindAddress
}
