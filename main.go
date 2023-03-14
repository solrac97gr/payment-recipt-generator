package main

import (
	"github.com/solrac97gr/payment-recipt-generator/factories"
)

func main() {
	configApplication := factories.NewConfigFileApplication("config.json")
	config, err := configApplication.GetConfig()
	if err != nil {
		panic(err)
	}

	_PDFGenerator := factories.NewPDFGeneratorApp("./out/", config)
	err = _PDFGenerator.GenerateInvoicePDF()
	if err != nil {
		panic(err)
	}

}
