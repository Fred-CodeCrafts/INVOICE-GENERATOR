# 📄 Invoice Generator API (Go)

This is a lightweight, modular, and testable Go microservice that generates professional PDF invoices based on dynamic user input.

Perfect for:
- Freelancers or businesses needing printable invoices
- Portfolio demonstration of Go backend skills
- Teaching resource for PDF generation, structuring, and APIs

---

## 🚀 Features

- Generate **PDF invoices** with:
  - Dynamic client and seller data
  - List of items/services
  - Tax, shipping, discounts, and total calculation
  - Custom notes, terms, and payment methods
- Clean and responsive invoice design
- Modular code structure (easy to expand/test)
- Designed with real-world invoice elements
- 🧪 Testable components with sample data

---



## 📦 Requirements

- Go 1.18 or newer
- gofpdf PDF library

Install dependencies:

go get github.com/jung-kurt/gofpdf

---

## 💡 Example Input (JSON)

{
  "invoice_number": "INV-00123",
  "purchase_order": "PO-456789",
  "issue_date": "2025-07-09",
  "due_date": "2025-07-16",
  "currency_symbol": "Rp",
  "seller": {
    "name": "My Company",
    "address": "Jl. Jakarta No. 1",
    "email": "contact@mycompany.com",
    "phone": "08123456789"
  },
  "customer": {
    "name": "Client Corp",
    "address": "Jl. Surabaya No. 2",
    "email": "client@corp.com",
    "phone": "08234567890"
  },
  "items": [
    {
      "description": "Web Development",
      "quantity": 1,
      "unit_price": 5000000,
      "tax_percent": 10
    },
    {
      "description": "Hosting (1 year)",
      "quantity": 1,
      "unit_price": 1000000,
      "tax_percent": 0
    }
  ],
  "tax_rate": 11,
  "shipping_cost": 0,
  "discount": 0,
  "notes": "Thanks for your business!",
  "terms": "Payment due within 7 days.",
  "payment_methods": "Bank Transfer (BCA, Mandiri)",
  "payment_terms": "Net 7"
}

---

## 🧪 Run the Project

go run cmd/main.go

Then call the API:

curl -X POST http://localhost:8080/generate \
  -H "Content-Type: application/json" \
  -d @data/customers.json

Or use Postman to test the /generate endpoint with JSON.

---

## 🖨️ Output

A PDF file will be saved under output/invoice_<timestamp>.pdf with a professional layout, ready to print or send.

---

## 📚 Future Improvements

- Web frontend (React/Next.js)
- PDF templates and theme switching
- QR code for payments
- Email invoice directly from API
- Authentication & user accounts

---

## 📄 License

MIT — Free to use, share, and improve.

---

## 🙌 Author

Made with 💻 and ☕ by FREDERICK GARNER WIBOWO(https://github.com/Fred-CodeCrafts)
