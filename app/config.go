package app

type Config interface {
	GetHTTPBindAddress() string
	GetGRPCBindAddress() string
}
type config struct {
	httpBindAddress string
	grpcBindAddress string
}

func NewConfig(httpAddr string, grpcAddr string) Config {
	return config{httpBindAddress: httpAddr, grpcBindAddress: grpcAddr}
}

func (c config) GetHTTPBindAddress() string {
	return c.httpBindAddress
}
func (c config) GetGRPCBindAddress() string {
	return c.grpcBindAddress
}
