package config

type database struct {
	Driver  string
	Address string
}

type config struct {
}

func GetConfig() *config {

	return new(config)
}

func (c *config) GetDatabase() database {
	return database{
		Driver:  `mysql`,
		Address: `root:fushihao@/skate?charset=utf8`,
	}
}

func (c *config) GetVersion() string {
	return "/v1"
}
