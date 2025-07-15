package invoice

import "errors"

func ValidateInvoice(inv Invoice) error {
	if inv.InvoiceNumber == "" {
		return errors.New("invoice number is required")
	}
	if inv.Customer.Name == "" || inv.Customer.Email == "" {
		return errors.New("customer name and email are required")
	}
	if len(inv.Items) == 0 {
		return errors.New("at least one item is required")
	}
	return nil
}
