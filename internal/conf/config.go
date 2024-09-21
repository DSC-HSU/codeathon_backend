package conf

import (
	"codeathon.runwayclub.dev/domain"
	"github.com/spf13/viper"
)

var Config *domain.Config

// ReadConfig
/*
 * Reads the configuration from the json file
 * @param jsonUrl string
 * @return *domain.Config
 * @return error
 */
func ReadConfig(jsonUrl string) error {
	// using Viper
	viper.SetConfigFile(jsonUrl)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	Config = &domain.Config{}
	err = viper.Unmarshal(Config)
	if err != nil {
		return err
	}
	return nil
}
