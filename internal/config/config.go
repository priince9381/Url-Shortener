package config

import "github.com/priince9381/Url-Shortener/app/internal/database"

type Config struct {
	App      App
	Database database.Database
}

type App struct {
	Env string
}
