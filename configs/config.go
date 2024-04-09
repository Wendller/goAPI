package configs

import (
	gorm_repositories "github.com/Wendller/goexpert/goAPI/internal/infra/database/repositories/gorm"
	"github.com/Wendller/goexpert/goAPI/internal/infra/web/handlers"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type cfg struct {
	DBDriver      string `mapstructure:"DATABASE_DRIVER"`
	DBHost        string `mapstructure:"DATABASE_HOST"`
	DBPort        string `mapstructure:"DATABASE_PORT"`
	DBUser        string `mapstructure:"DATABASE_USER"`
	DBPassword    string `mapstructure:"DATABASE_PASSWORD"`
	DBName        string `mapstructure:"DATABASE_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_POST"`
}

var config *cfg

func LoadConfig(path string) (*cfg, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config, nil
}

func InitializeRepositories(db *gorm.DB) *gorm_repositories.Repositories {
	return gorm_repositories.NewRepositories(db)
}

func InitializeHandlers(repositories *gorm_repositories.Repositories) *handlers.Handlers {
	return handlers.NewHandlers(repositories)
}
