package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/signintech/gopdf"
	"github.com/solrac97gr/payment-recipt-generator/pdf-generator/domain/ports"
)

// PDFGeneratorFolderRepository is a repository that saves the PDF file in a folder
type PDFGeneratorFolderRepository struct {
	filePath string
}

var _ ports.PDFGeneratorRepository = &PDFGeneratorFolderRepository{}

// NewPDFGeneratorFolderRepository returns a new instance of PDFGeneratorFolderRepository
func NewPDFGeneratorFolderRepository(filePath string) *PDFGeneratorFolderRepository {
	return &PDFGeneratorFolderRepository{filePath: filePath}
}

// errors
var (
	ErrWrittingPDF = errors.New("error writting pdf file")
)

// SavePDFFile saves the PDF file in a folder
func (r *PDFGeneratorFolderRepository) SavePDFFile(pdf gopdf.GoPdf) error {
	//Today date in string
	date := time.Now().Format("2006-01-02")

	route := fmt.Sprintf("%s%s-invoice.pdf", r.filePath, date)

	err := pdf.WritePdf(route)
	if err != nil {
		log.Println(err)
		return ErrWrittingPDF
	}

	return nil
}
