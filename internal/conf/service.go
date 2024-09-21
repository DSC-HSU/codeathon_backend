package conf

import (
	"codeathon.runwayclub.dev/domain"
	"context"
	"github.com/ServiceWeaver/weaver"
)

type ConfigService interface {
	GetConfig(ctx context.Context) (*domain.Config, error)
}

type configService struct {
	weaver.Implements[ConfigService]
}

func (c configService) GetConfig(ctx context.Context) (*domain.Config, error) {
	if Config == nil {
		// read config
		err := ReadConfig("./env/config.json")
		if err != nil {
			return nil, err
		}
	}
	// clone the config
	config := *Config
	return &config, nil
}
