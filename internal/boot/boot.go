package boot

import (
	"context"
	"log"
	"os"

	"github.com/priince938/app/internal/config"
	provider_config "github.com/priince938/app/internal/provider/config"
	config_reader "github.com/priince938/app/pkg/config"
)

func NewContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return ctx
}

func Init() {
	// Init config
	var conf config.Config

	err := config_reader.NewDefaultConfig().Load(GetEnv(), &conf)
	if err != nil {
		log.Fatal(err)
	}

	provider_config.InitializeConfig(conf)
}

func GetEnv() string {
	// Fetch env for bootstrapping
	environment := os.Getenv("APP_MODE")
	if environment == "" {
		environment = "dev"
	}

	return environment
}
