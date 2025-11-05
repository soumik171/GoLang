package main

import "fmt"

// Maintain open-close principle: code is open for extention but close for modification

type paymenter interface {
	pay(ammount float32)
	refund(amount float32, account string)
}

type payment struct {
	// Dependency inversion: now gateway is not dependent to a concreate gateway
	gateway paymenter
}

// add methods to struct
func (p payment) makePayment(ammount float32) {
	// If want to use razorpay, then have to create an instance
	// But if need multiple gateways, we need to modify, so build in such way that there's no need to change just add, so at that case, have to create instance inside the main

	// razorPaymentGW := razorpay{}
	// razorPaymentGW.pay(ammount)

	// just pass the value:
	p.gateway.pay(ammount)

}

// As we need payment gateway, so use another struct
// add razor payment gateway:

type razorpay struct {
}

// add methods to razor struct
func (r razorpay) pay(ammount float32) {
	// logic

	fmt.Println("making payment using razorpay", ammount)
}

// add stripe payment gateway:

type stripepay struct {
}

func (s stripepay) pay(ammount float32) {
	fmt.Println("making payment using stripepay", ammount)
}

// add paypal payment gateway:

type paypalpay struct {
}

func (p paypalpay) pay(ammount float32) {
	fmt.Println("making payment using paypalpay", ammount)
}

func (p paypalpay) refund(ammount float32, account string) {
}

func main() {

	// instance create of the gateway:

	// razorGW := razorpay{}
	// stripeGW := stripepay{}
	paypalGW := paypalpay{}

	// pass instance into makePayment
	newPayment := payment{
		gateway: paypalGW,
	}

	newPayment.makePayment(1500.00)

}
