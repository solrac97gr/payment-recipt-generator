package ports

type PDFGeneratorApplication interface {
	GenerateInvoicePDF() error
}
