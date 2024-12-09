package config

import (
	"sync"

	"github.com/spf13/viper"
)

type App struct {
	Name string `mapstructure:"APP_NAME"`
	Env  string `mapstructure:"APP_ENV"`
}

type Server struct {
	Port string `mapstructure:"SERVER_PORT"`
}

type Database struct {
	User string `mapstructure:"DATABASE_USER"`
	Pass string `mapstructure:"DATABASE_PASS"`
	Host string `mapstructure:"DATABASE_HOST"`
	Port int    `mapstructure:"DATABASE_PORT"`
	Name string `mapstructure:"DATABASE_NAME"`
}

type Jwt struct {
	Secret string `mapstructure:"JWT_SECRET"`
}

type Config struct {
	App      `mapstructure:",squash"`
	Server   `mapstructure:",squash"`
	Database `mapstructure:",squash"`
	Jwt      `mapstructure:",squash"`
}

var (
	once sync.Once
	cfg  *Config
)

func Load() {
	once.Do(func() {
		viper.AddConfigPath(".")
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	})
}

func Get() *Config {
	return cfg
}
