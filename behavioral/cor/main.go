package main

import "fmt"

type Account interface {
	Pay(amountToPay float32)
	SetNext(next Account)
}

type Bank struct {
	balance float32
	next    Account
}

var _ Account = &Bank{}

func NewBank(balance float32) *Bank {
	return &Bank{
		balance: balance,
	}
}

func (b *Bank) canPay(amount float32) bool {
	return b.balance >= amount
}

func (b *Bank) Pay(amountToPay float32) {
	if b.canPay(amountToPay) {
		b.balance -= amountToPay
		fmt.Printf("Paid %f using Bank!\n", amountToPay)
		return
	}
	fmt.Println("Cannot pay using Bank, Processing...")
	if b.next != nil {
		b.next.Pay(amountToPay)
		return
	}
	fmt.Println("Insufficient account balance.")
}

func (b *Bank) SetNext(next Account) {
	b.next = next
}

type Paypal struct {
	balance float32
	next    Account
}

var _ Account = &Paypal{}

func NewPaypal(balance float32) *Paypal {
	return &Paypal{
		balance: balance,
	}
}

func (b *Paypal) canPay(amount float32) bool {
	return b.balance >= amount
}

func (b *Paypal) Pay(amountToPay float32) {
	if b.canPay(amountToPay) {
		b.balance -= amountToPay
		fmt.Printf("Paid %f using Paypal!\n", amountToPay)
		return
	}
	fmt.Println("Cannot pay using Paypal, Processing...")
	if b.next != nil {
		b.next.Pay(amountToPay)
		return
	}
	fmt.Println("Insufficient account balance.")
}

func (b *Paypal) SetNext(next Account) {
	b.next = next
}

type Bitcoin struct {
	balance float32
	next    Account
}

var _ Account = &Bitcoin{}

func NewBitcoin(balance float32) *Bitcoin {
	return &Bitcoin{
		balance: balance,
	}
}

func (b *Bitcoin) canPay(amount float32) bool {
	return b.balance >= amount
}

func (b *Bitcoin) Pay(amountToPay float32) {
	if b.canPay(amountToPay) {
		b.balance -= amountToPay
		fmt.Printf("Paid %f using Bitcoin!\n", amountToPay)
		return
	}
	fmt.Println("Cannot pay using Bitcoin, Processing...")
	if b.next != nil {
		b.next.Pay(amountToPay)
		return
	}
	fmt.Println("Insufficient account balance.")
}

func (b *Bitcoin) SetNext(next Account) {
	b.next = next
}

func main() {
	// 准备一条如下所示的链
	//      bank -> paypal -> bitcoin
	//
	// 第一优先级是 bank
	//      如果 bank 无法支付，则使用 paypal
	//      如果 paypal 无法支付，则使用 bitcoin

	bank := NewBank(100)       // 余额为 100 的 Bank
	paypal := NewPaypal(200)   // 余额为 200 的 Paypal
	bitcoin := NewBitcoin(300) // 余额为 300 的 Bitcoin

	bank.SetNext(paypal)
	paypal.SetNext(bitcoin)

	bank.Pay(500)

	// 输出结果
	// ==============
	// Cannot pay using Bank, Processing...
	// Cannot pay using Paypal, Processing...
	// Paid 259.000000 using Bitcoin!
}
