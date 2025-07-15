package invoice

type Customer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	TaxID   string `json:"tax_id"`
}

type Item struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TaxPercent  float64 `json:"tax_percent"`
}

type Invoice struct {
	InvoiceNumber  string
	IssueDate      string
	DueDate        string
	PurchaseOrder  string
	ShippingInfo   string
	CurrencyCode   string
	CurrencySymbol string
	Seller         Customer
	Customer       Customer
	Items          []Item
	ShippingCost   float64
	TaxRate        float64
	Discount       float64
	Status         string
	Notes          string
	Terms          string
	GeneratedAt    string
	PaymentMethods string `json:"payment_methods"`
	PaymentTerms   string `json:"payment_terms"`
}
