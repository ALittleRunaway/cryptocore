package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	ServiceName = "thourus.cryptocore"
	EnvPrefix   = "THOURUS_CRYPTOCORE"
)

type Config struct {
	Log    *LogConfig
	Grpc   *GrpcConfig
	Crypto *CryptoConfig
}

func InitError(err error) {
	if err != nil {
		panic(err)
	}
}

// InitConfig initialise all the configurations
func InitConfig() *Config {
	return &Config{
		Log:    InitLogConfig(),
		Grpc:   InitGrpcConfig(),
		Crypto: InitCryptoConfig(),
	}
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("There is no .env file provided!")
	}
	viper.SetDefault("cryptocore.dedicated_env", "DEV")
	InitError(viper.BindEnv("cryptocore.dedicated_env", EnvPrefix+"_DEDICATED_ENV"))
}
