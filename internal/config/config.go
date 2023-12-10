package config

import "github.com/priince938/app/internal/database"

type Config struct {
	App      App
	Database database.Database
}

type App struct {
	Env string
}
