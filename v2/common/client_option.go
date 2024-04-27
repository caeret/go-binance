package common

type ClientConfig struct {
	UseTestnet bool
}

func ParseClientConfig(opts ...ClientOptionFunc) ClientConfig {
	cfg := ClientConfig{UseTestnet: false}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

type ClientOptionFunc func(*ClientConfig)

func UseTestnet(useTestnet bool) ClientOptionFunc {
	return func(cfg *ClientConfig) {
		cfg.UseTestnet = useTestnet
	}
}
