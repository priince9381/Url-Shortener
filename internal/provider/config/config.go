package config

import "github.com/priince938/app/internal/config"

var (
	Config config.Config
)

func InitializeConfig(client config.Config) {
	Config = client
}
