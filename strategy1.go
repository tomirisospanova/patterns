package main

import "fmt"

type PaymentStrategy interface {
	Pay() error
}

type PaymentContext struct {
	paymentStrategy PaymentStrategy
}

func (pc *PaymentContext) processOrder(product string) {
	err := pc.paymentStrategy.Pay()
	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}
	fmt.Println("Payment successful for product:", product)
}

type CardPayment struct {
	cardNumber, cvv string
}

func (cp *CardPayment) Pay() error {
	fmt.Println("Processing card payment...")
	return nil
}

func NewCardPaymentStrategy(cardNumber, cvv string) PaymentStrategy {
	return &CardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

type KaspiQRPayment struct{}

func (pp *KaspiQRPayment) Pay() error {
	fmt.Println("Processing KaspiQR payment...")
	return nil
}

func NewKaspiQRPaymentStrategy() PaymentStrategy {
	return &KaspiQRPayment{}
}

type KaspiPayment struct{}

func (qp *KaspiPayment) Pay() error {
	fmt.Println("Processing Kaspi payment...")
	return nil
}

func NewKaspiPaymentStrategy() PaymentStrategy {
	return &KaspiPayment{}
}

func main() {
	product := "computer"
	payWay := 3

	var paymentStrategy PaymentStrategy
	switch payWay {
	case 1:
		paymentStrategy = NewCardPaymentStrategy("12345", "12345")
	case 2:
		paymentStrategy = NewKaspiQRPaymentStrategy()
	case 3:
		paymentStrategy = NewKaspiPaymentStrategy()
	default:
		fmt.Println("Invalid payment option")
		return
	}

	paymentContext := PaymentContext{paymentStrategy}
	paymentContext.processOrder(product)
}
