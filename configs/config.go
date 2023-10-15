package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	SmsSvcUrl  string `mapstructure:"SMS_SVC_URL"`
	RedUrl     string `mapstructure:"RED_URL"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbSslMode  string `mapstructure:"DB_SSLMODE"`
}

func InitConfig() (cfg Config) {
	viper.SetConfigFile("./configs/envs/dev.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read config: " + err.Error())
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("failed to unmarshal config: " + err.Error())
	}

	return
}
