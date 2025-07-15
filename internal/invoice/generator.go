package invoice

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func GenerateInvoicePDF(inv Invoice, filepath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTitle("Invoice "+inv.InvoiceNumber, false)
	pdf.AddPage()
	pdf.SetMargins(15, 15, 15)

	// Title & Date
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(100, 10, "INVOICE")

	pdf.SetFont("Arial", "", 10)
	pdf.SetXY(160, 15)
	pdf.CellFormat(40, 10, fmt.Sprintf("Generated on \n%s", inv.IssueDate), "", 1, "R", false, 0, "")

	// DRAFT badge
	pdf.SetFillColor(200, 200, 200)
	pdf.SetFont("Arial", "B", 9)
	pdf.SetXY(15, 25)
	pdf.CellFormat(20, 6, "DRAFT", "", 1, "C", true, 0, "")

	pdf.Ln(5)

	// FROM / TO / DETAILS
	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(60, 6, "From:")
	pdf.Cell(60, 6, "To:")
	pdf.Cell(60, 6, "Invoice Details:")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(60, 5, inv.Seller.Name)
	pdf.Cell(60, 5, inv.Customer.Name)
	pdf.Cell(60, 5, fmt.Sprintf("Invoice #: %s", inv.InvoiceNumber))
	pdf.Ln(5)
	pdf.Cell(60, 5, inv.Seller.Email)
	pdf.Cell(60, 5, inv.Customer.Email)
	pdf.Cell(60, 5, fmt.Sprintf("Date: %s", inv.IssueDate))
	pdf.Ln(5)
	pdf.Cell(60, 5, inv.Seller.Phone)
	pdf.Cell(60, 5, inv.Customer.Phone)
	pdf.Cell(60, 5, fmt.Sprintf("Due Date: %s", inv.DueDate))
	pdf.Ln(5)
	pdf.Cell(60, 5, inv.Seller.Address)
	pdf.Cell(60, 5, inv.Customer.Address)
	pdf.Cell(60, 5, fmt.Sprintf("Terms: %s", inv.Terms))
	pdf.Ln(10)

	// Table Header
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(245, 245, 245)

	pdf.CellFormat(90, 8, "Description", "1", 0, "", true, 0, "")
	pdf.CellFormat(20, 8, "Qty", "1", 0, "C", true, 0, "")
	pdf.CellFormat(35, 8, "Unit Price", "1", 0, "R", true, 0, "")
	pdf.CellFormat(35, 8, "Amount", "1", 1, "R", true, 0, "")

	// Table Body
	pdf.SetFont("Arial", "", 10)
	subtotal := 0.0
	for _, item := range inv.Items {
		amount := float64(item.Quantity) * item.UnitPrice
		subtotal += amount

		pdf.CellFormat(90, 8, item.Description, "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 8, fmt.Sprintf("%d", item.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(35, 8, fmt.Sprintf("%s%.2f", inv.CurrencySymbol, item.UnitPrice), "1", 0, "R", false, 0, "")
		pdf.CellFormat(35, 8, fmt.Sprintf("%s%.2f", inv.CurrencySymbol, amount), "1", 1, "R", false, 0, "")
	}

	// Totals Section
	pdf.Ln(4)
	pdf.Cell(120, 6, "")
	pdf.CellFormat(30, 6, "Subtotal:", "", 0, "R", false, 0, "")
	pdf.CellFormat(30, 6, fmt.Sprintf("%s%.2f", inv.CurrencySymbol, subtotal), "", 1, "R", false, 0, "")

	tax := subtotal * inv.TaxRate / 100
	total := subtotal + tax

	pdf.Cell(120, 6, "")
	pdf.CellFormat(30, 6, fmt.Sprintf("Tax (%.1f%%):", inv.TaxRate), "", 0, "R", false, 0, "")
	pdf.CellFormat(30, 6, fmt.Sprintf("%s%.2f", inv.CurrencySymbol, tax), "", 1, "R", false, 0, "")
	// Match font size with the rest (10pt)
	pdf.SetFont("Arial", "B", 10)

	// Prepare label and value
	label := "Total Amount Due:"
	value := fmt.Sprintf("%s%.2f", inv.CurrencySymbol, total)

	// Dynamically calculate widths
	labelWidth := pdf.GetStringWidth(label) + 4
	valueWidth := pdf.GetStringWidth(value) + 6

	// Ensure minimum width
	if labelWidth < 40 {
		labelWidth = 24
	}
	if valueWidth < 40 {
		valueWidth = 40
	}

	totalBoxWidth := labelWidth + valueWidth
	pageWidth, _ := pdf.GetPageSize()
	marginLeft, _, _, _ := pdf.GetMargins()
	x := pageWidth - marginLeft - totalBoxWidth

	// Position
	pdf.SetXY(x, pdf.GetY())

	// Label (blue, no border, no fill)
	pdf.SetTextColor(0, 102, 255)
	pdf.CellFormat(labelWidth, 6, label, "", 0, "L", false, 0, "")

	// Value (black, no border, no fill)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(valueWidth, 6, value, "", 1, "R", false, 0, "")

	// Horizontal line after
	pdf.Ln(4)
	pdf.Line(marginLeft, pdf.GetY(), pageWidth-marginLeft, pdf.GetY())
	pdf.Ln(4)

	// Payment Info
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(95, 6, "Payment Methods:")
	pdf.Cell(95, 6, "Payment Terms:")
	pdf.Ln(6)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(95, 6, inv.PaymentMethods)
	pdf.Cell(95, 6, inv.PaymentTerms)
	pdf.Ln(8)

	// Notes
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(30, 6, "Notes")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(255, 255, 200)
	pdf.MultiCell(0, 6, inv.Notes, "0", "", true)
	pdf.Ln(5)

	// Footer
	pdf.SetFont("Arial", "I", 8)
	pdf.SetTextColor(100, 100, 100)
	pdf.Cell(0, 5, "This invoice was generated electronically and is valid without signature.")
	pdf.Ln(4)
	pdf.Cell(0, 5, fmt.Sprintf("Generated on %s", inv.IssueDate))

	os.MkdirAll("output", os.ModePerm)
	return pdf.OutputFileAndClose(filepath)
}
