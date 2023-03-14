package application

import (
	"fmt"
	"log"
	"time"

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
	// Invoice Date
	DrawInvoiceDate(&pdf)
	// Due Date
	DrawDueDate(&pdf)
	// Items Table
	DrawItemsTable(&pdf, &a.config.PaymentDetails.Work, &a.config.PaymentDetails.Amount)
	// Invoice Summary
	DrawInvoiceSummary(&pdf, &a.config.PaymentDetails.Amount)

	err = a.repo.SavePDFFile(pdf)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}

func DrawHeader(pdf *gopdf.GoPdf) {
	pdf.SetFont("roboto-light", "", 10)
	pdf.SetTextColor(78, 78, 78)
	pdf.SetXY(300, 20)
	pdf.Cell(nil, "Invoice")
}

func DrawIssuerInformation(pdf *gopdf.GoPdf, issuer *models.Issuer) {
	horizontalStart := 20.0
	verticalStart := 70.0
	separator := 20.0
	fontSize := 9.0
	pdf.SetTextColor(78, 78, 78)

	pdf.SetFont("roboto-light", "", fontSize)
	pdf.SetFontSize(8)
	pdf.SetXY(horizontalStart, verticalStart)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, "FROM")
	pdf.SetXY(horizontalStart, verticalStart+separator*2)
	pdf.SetFont("roboto-bold", "", fontSize)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, issuer.Name)
	pdf.SetTextColor(78, 78, 78)
	pdf.SetFont("roboto-regular", "", fontSize)
	pdf.SetXY(horizontalStart, verticalStart+separator*3)
	pdf.Cell(nil, issuer.Address)
	pdf.SetXY(horizontalStart, verticalStart+separator*4)
	pdf.Cell(nil, issuer.Email)
	pdf.SetXY(horizontalStart, verticalStart+separator*5)
	pdf.Cell(nil, issuer.Phone)
	pdf.SetXY(horizontalStart, verticalStart+separator*6)
	pdf.Cell(nil, issuer.Website)
}

func DrawCompanyInformation(pdf *gopdf.GoPdf, company *models.Company) {
	horizontalStart := 300.0
	verticalStart := 70.0
	separator := 20.0
	fontSize := 9.0
	pdf.SetTextColor(78, 78, 78)

	pdf.SetFont("roboto-light", "", fontSize)
	pdf.SetFontSize(8)
	pdf.SetXY(horizontalStart, verticalStart)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, "TO")
	pdf.SetXY(horizontalStart, verticalStart+separator*2)
	pdf.SetFont("roboto-bold", "", fontSize)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, company.Name)
	pdf.SetTextColor(78, 78, 78)
	pdf.SetFont("roboto-regular", "", fontSize)
	pdf.SetXY(horizontalStart, verticalStart+separator*3)
	pdf.Cell(nil, company.Representative)
	pdf.SetXY(horizontalStart, verticalStart+separator*4)
	pdf.Cell(nil, company.Address)
	pdf.SetXY(horizontalStart, verticalStart+separator*5)
	pdf.Cell(nil, company.Email)
	pdf.SetXY(horizontalStart, verticalStart+separator*6)
	pdf.Cell(nil, company.Website)
}

func DrawInvoiceDate(pdf *gopdf.GoPdf) {
	horizontalStart := 20.0
	verticalStart := 250.0

	pdf.SetFont("roboto-black", "", 9)
	pdf.SetXY(horizontalStart, verticalStart)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, "Invoice Date:")
	pdf.SetTextColor(78, 78, 78)
	pdf.SetXY(horizontalStart+55, verticalStart)
	pdf.SetFont("roboto-regular", "", 9)
	pdf.Cell(nil, time.Now().Format("02/01/2006"))
}

func DrawDueDate(pdf *gopdf.GoPdf) {
	horizontalStart := 300.0
	verticalStart := 250.0

	pdf.SetFont("roboto-black", "", 9)
	pdf.SetXY(horizontalStart, verticalStart)
	pdf.SetTextColor(0, 0, 0)
	pdf.Cell(nil, "Due Date:")
	pdf.SetTextColor(78, 78, 78)
	pdf.SetXY(horizontalStart+45, verticalStart)
	pdf.SetFont("roboto-regular", "", 9)
	pdf.Cell(nil, time.Now().Format("02/01/2006"))
}

func DrawItemsTable(pdf *gopdf.GoPdf, work *models.Work, amount *models.Amount) {
	DrawItemsTableHeader(pdf)
	horizontalStart := 30.0
	verticalStart := 330.0
	separator := 80.0
	fontSize := 9.0

	pdf.SetTextColor(78, 78, 78)
	pdf.SetFont("roboto-regular", "", fontSize)
	pdf.SetXY(horizontalStart, verticalStart)
	pdf.Cell(nil, work.Description)
	pdf.SetXY(horizontalStart+separator*2, verticalStart)
	pdf.Cell(nil, "1")
	pdf.SetXY(horizontalStart+separator*4, verticalStart)
	pdf.Cell(nil, amount.Currency)
	pdf.SetXY(horizontalStart+separator*5, verticalStart)
	totalString := fmt.Sprintf("%v", amount.Total)
	pdf.Cell(nil, totalString)

	DrawItemsTableFooter(pdf, amount)
}

func DrawItemsTableHeader(pdf *gopdf.GoPdf) {
	pdf.SetLineWidth(1)
	pdf.Line(20, 300, 565, 300)
	horizontalStart := 30.0
	verticalStart := 306.0
	separator := 80.0
	fontSize := 9.0
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("roboto-black", "", fontSize)
	pdf.SetXY(horizontalStart, verticalStart)
	pdf.Cell(nil, "Item")
	pdf.SetXY(horizontalStart+separator*2, verticalStart)
	pdf.Cell(nil, "Quantity")
	pdf.SetXY(horizontalStart+separator*4, verticalStart)
	pdf.Cell(nil, "Currency")
	pdf.SetXY(horizontalStart+separator*5, verticalStart)
	pdf.Cell(nil, "Total")

	pdf.SetLineWidth(1)
	pdf.Line(20, 320, 565, 320)
}

func DrawItemsTableFooter(pdf *gopdf.GoPdf, amount *models.Amount) {
	pdf.SetLineWidth(1)
	pdf.Line(20, 360, 565, 360)
}

func DrawInvoiceSummary(pdf *gopdf.GoPdf, amount *models.Amount) {
	horizontalStart := 30.0
	verticalStart := 410.0
	separator := 80.0
	fontSize := 9.0

	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("roboto-black", "", fontSize)
	pdf.SetXY(horizontalStart+separator*4, verticalStart)
	pdf.Cell(nil, "Currency")
	pdf.SetXY(horizontalStart+separator*5, verticalStart)
	pdf.Cell(nil, amount.Currency)

	pdf.SetFont("roboto-black", "", fontSize)
	pdf.SetXY(horizontalStart+separator*4, verticalStart+20)
	pdf.Cell(nil, "Total")
	pdf.SetXY(horizontalStart+separator*5, verticalStart+20)
	totalString := fmt.Sprintf("%v", amount.Total)
	pdf.Cell(nil, totalString)
}

func AddFonts(pdf *gopdf.GoPdf) error {

	err := pdf.AddTTFFont("roboto-black", "./font/Roboto-Black.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-bold", "./font/Roboto-Bold.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-light", "./font/Roboto-Light.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	err = pdf.AddTTFFont("roboto-regular", "./font/Roboto-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}
