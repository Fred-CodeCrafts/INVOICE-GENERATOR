package store

import (
	"encoding/json"
	"invoice-generator/internal/invoice"
	"io/ioutil"
	"os"
)

type FileStore struct {
	path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{path: path}
}

func (fs *FileStore) SaveCustomer(c invoice.Customer) error {
	existing, _ := fs.GetCustomers()
	existing = append(existing, c)
	data, _ := json.MarshalIndent(existing, "", "  ")
	return ioutil.WriteFile(fs.path, data, 0644)
}

func (fs *FileStore) GetCustomers() ([]invoice.Customer, error) {
	if _, err := os.Stat(fs.path); os.IsNotExist(err) {
		return []invoice.Customer{}, nil
	}
	data, err := ioutil.ReadFile(fs.path)
	if err != nil {
		return nil, err
	}
	var customers []invoice.Customer
	err = json.Unmarshal(data, &customers)
	return customers, err
}
