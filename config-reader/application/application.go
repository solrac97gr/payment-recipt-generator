package application

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/solrac97gr/payment-recipt-generator/config-reader/domain/models"
	"github.com/solrac97gr/payment-recipt-generator/config-reader/domain/ports"
)

type ConfigApplication struct {
	repo ports.ConfigRepository
}

var _ ports.ConfigApp = &ConfigApplication{}

// NewConfigApplication returns a new instance of ConfigApplication
func NewConfigApplication(repo ports.ConfigRepository) *ConfigApplication {
	return &ConfigApplication{
		repo: repo,
	}
}

// error messages
var (
	// ErrConfigFileNotFound is returned when the config file is not found
	ErrConfigFileNotFound = errors.New("config file not found")
	// ErrUnmarshalingConfig is returned when the config file can't be unmarshaled
	ErrUnmarshalingConfig = errors.New("error unmarshaling config file")
)

// GetConfig returns the configuration of the application
func (c *ConfigApplication) GetConfig() (*models.Config, error) {
	file, err := c.repo.GetConfigFile()
	if err != nil {
		log.Println(err)
		return nil, ErrConfigFileNotFound
	}

	var config models.Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println(err)
		return nil, ErrUnmarshalingConfig
	}

	return &config, nil
}
