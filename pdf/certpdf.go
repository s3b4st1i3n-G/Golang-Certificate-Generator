package pdf

import (
	"fmt"
	"gencert/cert"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputDir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return p, err
	}
	p = &PdfSaver{
		OutputDir: outputDir,
	}
	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// header
	header(pdf, &cert)
	pdf.Ln(30)

	// body
	body(pdf, &cert)

	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}
	fmt.Sprintf("Saved certificate to '%v'\n", path)
	return nil
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func body(pdf *gofpdf.Fpdf, c *cert.Cert) {
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelPresented, "C")
	pdf.Ln(30)

	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, c.Name, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelParticipation, "C")
	pdf.Ln(30)

	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, c.LabelDate, "C")
}
