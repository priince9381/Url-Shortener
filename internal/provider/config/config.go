package config

import "github.com/priince9381/Url-Shortener/app/internal/config"

var (
	Config config.Config
)

func InitializeConfig(client config.Config) {
	Config = client
}
