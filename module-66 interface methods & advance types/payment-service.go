package main

import "fmt"

type Bkash struct {
	apiKey string
}

func (bk *Bkash) pay(amount float64) {
	fmt.Printf("Paying %f tk with Bkash\n", amount)
}

type Nagad struct {
	apiKey string
}

func (ng *Nagad) pay(amount float64) {
	fmt.Printf("Paying %f tk with Nagad\n", amount)
}

type PaymentMethod interface {
	pay(amount float64)
}

type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) PaymentService {
	return PaymentService{method: method}
}

func (ps PaymentService) checkout() {
	ps.method.pay(100.00)
}

// development time testing purpose
type MockPaymentMethod struct{}

func (mockPM MockPaymentMethod) pay(amount float64) {
	fmt.Println("Testing payment method")
}

func payment_service_func() {
	bk := Bkash{apiKey: "xyz123"}
	// paymentService := PaymentService{} // method is required
	bkPaymentService := NewPaymentService(&bk)
	bkPaymentService.checkout()

	ng := Nagad{apiKey: "xyz123"}
	ngPaymentService := NewPaymentService(&ng)
	ngPaymentService.checkout()

	mockPm := MockPaymentMethod{}
	mockPaymentService := NewPaymentService(mockPm)
	mockPaymentService.checkout()
}
