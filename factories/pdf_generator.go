package factories

import (
	"github.com/solrac97gr/payment-recipt-generator/config-reader/domain/models"
	"github.com/solrac97gr/payment-recipt-generator/pdf-generator/application"
	"github.com/solrac97gr/payment-recipt-generator/pdf-generator/infrastructure/repository"
)

func NewPDFGeneratorApp(outputFolderPath string, config *models.Config) *application.PDFGeneratorApp {
	PDFGeneratorRepository := repository.NewPDFGeneratorFolderRepository(outputFolderPath)
	return application.NewPDFGeneratorApp(PDFGeneratorRepository, config)
}
