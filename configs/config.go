package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type cfg struct {
	DBDriver      string `mapstructure:"DATABASE_DRIVER"`
	DBHost        string `mapstructure:"DATABASE_HOST"`
	DBPort        string `mapstructure:"DATABASE_PORT"`
	DBUser        string `mapstructure:"DATABASE_USER"`
	DBPassword    string `mapstructure:"DATABASE_PASSWORD"`
	DBName        string `mapstructure:"DATABASE_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_POST"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
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

	config.TokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	return config, nil
}
