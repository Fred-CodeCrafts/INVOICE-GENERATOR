package store

import "invoice-generator/internal/invoice"

type Store interface {
	SaveCustomer(c invoice.Customer) error
	GetCustomers() ([]invoice.Customer, error)
}
