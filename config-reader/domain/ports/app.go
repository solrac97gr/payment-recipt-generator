package ports

import "github.com/solrac97gr/payment-recipt-generator/config-reader/domain/models"

type ConfigApp interface {
	GetConfig() (*models.Config, error)
}
