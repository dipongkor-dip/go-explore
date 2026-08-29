package mocPayment

import "fmt"

type MockPaymentMethod struct{}

func (mcp MockPaymentMethod) Pay(amount float64) {
	fmt.Println("Testing payment method")
}
