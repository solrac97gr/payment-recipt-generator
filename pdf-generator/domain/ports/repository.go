package ports

import "github.com/signintech/gopdf"

type PDFGeneratorRepository interface {
	SavePDFFile(gopdf.GoPdf) error
}
