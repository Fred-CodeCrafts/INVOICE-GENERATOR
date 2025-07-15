package invoice

import "testing"

func TestGenerateInvoicePDF(t *testing.T) {
	inv := Invoice{
		InvoiceNumber:  "TEST-001",
		IssueDate:      "2025-07-09",
		DueDate:        "2025-07-16",
		Customer:       Customer{Name: "Test User", Email: "test@example.com"},
		Seller:         Customer{Name: "Test Seller"},
		Items:          []Item{{"Item 1", 2, 50.0, 10}},
		CurrencySymbol: "$",
		TaxRate:        10,
		ShippingCost:   5,
		Discount:       10,
		Notes:          "Test invoice only.",
		Terms:          "Net 7 days",
		GeneratedAt:    "2025-07-09T00:00:00Z",
	}

	err := GenerateInvoicePDF(inv, "../../output/test_invoice.pdf")
	if err != nil {
		t.Errorf("Failed to generate PDF: %v", err)
	}
}
