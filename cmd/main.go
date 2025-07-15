package main

import (
	"invoice-generator/internal/invoice"
	"invoice-generator/internal/store"
	"log"
)

func main() {
	// Create a file-based customer store
	customerStore := store.NewFileStore("data/customers.json")

	// Load existing customers
	customers, _ := customerStore.GetCustomers()
	if len(customers) > 0 {
		log.Println("Using existing customer:", customers[0].Name)
	} else {
		log.Println("No customers found, using fallback")
	}

	customer := invoice.Customer{
		Name:    "John Doe",
		Address: "123 Main Street",
		Email:   "john@example.com",
		Phone:   "+62 812 3456 7890",
		TaxID:   "1234567890",
	}
	customerStore.SaveCustomer(customer)

	inv := invoice.Invoice{
		InvoiceNumber:  "INV-0001",
		IssueDate:      "2025-07-09",
		DueDate:        "2025-07-16",
		PurchaseOrder:  "PO-2025-7890",
		ShippingInfo:   "Delivery by JNE Express",
		CurrencyCode:   "IDR",
		CurrencySymbol: "Rp",
		Customer:       customer,
		Seller: invoice.Customer{
			Name:    "My Company",
			Address: "Jakarta, Indonesia",
			Email:   "sales@mycompany.com",
			Phone:   "+62 812 9876 5432",
			TaxID:   "VAT-1122334455",
		},

		Items: []invoice.Item{
			{
				Description: "Web Development",
				Quantity:    1,
				UnitPrice:   5000000,
				TaxPercent:  0,
			},
			{
				Description: "Hosting (1 year)",
				Quantity:    1,
				UnitPrice:   1000000,
				TaxPercent:  0,
			},
		},

		ShippingCost:   25000,
		TaxRate:        11,
		Discount:       50000,
		Status:         "Unpaid",
		Notes:          "Payment due within 7 days.",
		Terms:          "Late payments incur a 5% fee.",
		GeneratedAt:    "2025-07-09T14:00:00+07:00",
		PaymentMethods: "Bank Transfer, Credit Card, PayPal",
		PaymentTerms:   "Payment due within 30 days",
	}

	err := invoice.GenerateInvoicePDF(inv, "output/invoice_sample.pdf")
	if err != nil {
		log.Fatalf("Failed to generate invoice: %v", err)
	}
	log.Println("Invoice generated at output/invoice_sample.pdf")
}
