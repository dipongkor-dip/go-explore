package paymentService

import "fmt"

type Bkash struct {
	APIKey string
}

func (bk *Bkash) Pay(amount float64) {
	fmt.Printf("Paying %f tk with Bkash\n", amount)
}

type Nagad struct {
	APIKey string
}

func (ng *Nagad) Pay(amount float64) {
	fmt.Printf("Paying %f tk with Nagad\n", amount)
}

type paymentMethod interface {
	Pay(amount float64)
}

type paymentService struct {
	method paymentMethod
}

func NewPaymentService(method paymentMethod) *paymentService {
	return &paymentService{method: method}
}

func (ps paymentService) Checkout() {
	ps.method.Pay(100.00)
}
