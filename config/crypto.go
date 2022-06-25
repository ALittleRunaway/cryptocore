package config

import (
	"github.com/spf13/viper"
)

type CryptoConfig struct {
	SecretString string
}

func InitCryptoConfig() *CryptoConfig {

	cryptoConfig := CryptoConfig{
		SecretString: viper.GetString("cryptocore.secret_string"),
	}
	return &cryptoConfig
}

func init() {
	InitError(viper.BindEnv("cryptocore.secret_string", EnvPrefix+"_CRYPTO_SECRET_STRING"))
}
