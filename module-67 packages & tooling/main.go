package main

import (
	"fmt"
	"learn-package-module/mocPayment"
	"learn-package-module/paymentService"
	"learn-package-module/utility"
)

// payment_service "learn-package-module/payment-service"

func main() {
	type_handle_func()
	variadic_func()

	// different package handle
	// function name first latter must be a capital letter
	mul := utility.Multiply(5, 3)
	fmt.Println(mul)

	div := utility.Divide(5, 3)
	fmt.Println(div)

	bk := paymentService.Bkash{APIKey: "xyz123"}
	bkPaymentService := paymentService.NewPaymentService(&bk)
	bkPaymentService.Checkout()

	ng := paymentService.Nagad{APIKey: "xyz123"}
	ngPaymentService := paymentService.NewPaymentService(&ng)
	ngPaymentService.Checkout()

	mockPm := mocPayment.MockPaymentMethod{}
	mockPaymentService := paymentService.NewPaymentService(mockPm)
	mockPaymentService.Checkout()

}

// module = bunch of packages
// go run .
