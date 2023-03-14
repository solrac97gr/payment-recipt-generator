package repository

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/solrac97gr/payment-recipt-generator/config-reader/domain/ports"
)

type ConfigFileRepository struct {
	filePath string
}

var _ ports.ConfigRepository = &ConfigFileRepository{}

// NewConfigFileRepository returns a new instance of ConfigFileRepository
func NewConfigFileRepository(filePath string) *ConfigFileRepository {
	return &ConfigFileRepository{
		filePath: filePath,
	}
}

// errors
var (
	// ErrConfigFileCannotBeRead is returned when the config file cannot be read
	ErrConfigFileCannotBeRead = errors.New("config file cannot be read")
)

// GetConfigFile returns the content of the config file
func (c *ConfigFileRepository) GetConfigFile() ([]byte, error) {
	data, err := ioutil.ReadFile(c.filePath)
	if err != nil {
		log.Println(err)
		return nil, ErrConfigFileCannotBeRead
	}
	return data, nil
}
