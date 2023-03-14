package application

import (
	"log"

	"github.com/signintech/gopdf"
	"github.com/solrac97gr/payment-recipt-generator/config-reader/domain/models"
	"github.com/solrac97gr/payment-recipt-generator/pdf-generator/domain/ports"
)

// PDFGeneratorApp is the application that generates the PDF file
type PDFGeneratorApp struct {
	config *models.Config
	repo   ports.PDFGeneratorRepository
}

var _ ports.PDFGeneratorApplication = &PDFGeneratorApp{}

// NewPDFGeneratorApp returns a new instance of PDFGeneratorApp
func NewPDFGeneratorApp(repo ports.PDFGeneratorRepository, config *models.Config) *PDFGeneratorApp {
	return &PDFGeneratorApp{repo: repo, config: config}
}

// GenerateInvoicePDF generates the PDF file
func (a *PDFGeneratorApp) GenerateInvoicePDF() error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err := AddFonts(&pdf)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	// Header
	DrawHeader(&pdf)

	// From
	DrawIssuerInformation(&pdf, &a.config.PaymentDetails.Issuer)

	// To
	DrawCompanyInformation(&pdf, &a.config.PaymentDetails.Company)

	err = a.repo.SavePDFFile(pdf)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}

func DrawHeader(pdf *gopdf.GoPdf) {
	pdf.SetFont("roboto-medium", "", 14)
	pdf.SetXY(500, 20)
	pdf.Cell(nil, "Invoice")
}

func DrawIssuerInformation(pdf *gopdf.GoPdf, issuer *models.Issuer) {
	horizontalStart := 20.0
	fontSize := 9.0

	pdf.SetFont("roboto-light", "", fontSize)
	pdf.SetFontSize(8)
	pdf.SetXY(horizontalStart, 70)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, "FROM")
	pdf.SetXY(horizontalStart, 110)
	pdf.SetFont("roboto-bold", "", fontSize)
	pdf.Cell(nil, issuer.Name)
	pdf.SetFont("roboto-regular", "", fontSize)
	pdf.SetXY(horizontalStart, 130)
	pdf.Cell(nil, issuer.Address)
	pdf.SetXY(horizontalStart, 150)
	pdf.Cell(nil, issuer.Email)
	pdf.SetXY(horizontalStart, 170)
	pdf.Cell(nil, issuer.Phone)
	pdf.SetXY(horizontalStart, 190)
	pdf.Cell(nil, issuer.Website)
}

func DrawCompanyInformation(pdf *gopdf.GoPdf, company *models.Company) {
	horizontalStart := 300.0
	fontSize := 9.0

	pdf.SetFont("roboto-light", "", fontSize)
	pdf.SetFontSize(8)
	pdf.SetXY(horizontalStart, 70)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, "TO")
	pdf.SetXY(horizontalStart, 110)
	pdf.SetFont("roboto-bold", "", fontSize)
	pdf.Cell(nil, company.Name)
	pdf.SetFont("roboto-regular", "", fontSize)
	pdf.SetXY(horizontalStart, 130)
	pdf.Cell(nil, company.Representative)
	pdf.SetXY(horizontalStart, 150)
	pdf.Cell(nil, company.Address)
	pdf.SetXY(horizontalStart, 170)
	pdf.Cell(nil, company.Email)
	pdf.SetXY(horizontalStart, 190)
	pdf.Cell(nil, company.Website)
}

func AddFonts(pdf *gopdf.GoPdf) error {

	err := pdf.AddTTFFont("roboto-black", "./font/Roboto-Black.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-black-italic", "./font/Roboto-BlackItalic.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-bold", "./font/Roboto-Bold.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-bold-italic", "./font/Roboto-BoldItalic.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-italic", "./font/Roboto-Italic.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-light", "./font/Roboto-Light.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-light-italic", "./font/Roboto-LightItalic.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-medium", "./font/Roboto-Medium.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-medium-italic", "./font/Roboto-MediumItalic.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-regular", "./font/Roboto-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-thin", "./font/Roboto-Thin.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-thin-italic", "./font/Roboto-ThinItalic.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
