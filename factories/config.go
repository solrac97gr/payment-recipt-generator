package factories

import (
	"github.com/solrac97gr/payment-recipt-generator/config-reader/application"
	"github.com/solrac97gr/payment-recipt-generator/config-reader/infrastructure/repository"
)

func NewConfigFileApplication(filePath string) *application.ConfigApplication {
	configRepository := repository.NewConfigFileRepository(filePath)
	configApplication := application.NewConfigApplication(configRepository)
	return configApplication
}
