package httpServer



type config struct {
	host		string
	port		uint
}

func NewConfig(host string, port uint) *config {
	return &config{host: host, port: port}
}

func (c *config) Host() string {
	return c.host
}

func (c *config) Port() uint {
	return c.port
}


