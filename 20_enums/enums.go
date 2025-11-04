package main

import "fmt"

// enumerated types:conbined of custome types+const
// No need to pass value manually, increase readibilty

// Enum in int:For int inside const, if declare custom types for the first one, next one are reclared as a same type

type OrderStatus int // custom type

const (
	Received OrderStatus = iota // iota:untyped integer, that increments automatically
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

// enum in string:declare custom types for every line

type OrderStatusStr string

const (
	Receive OrderStatusStr = "received"
	Confirm OrderStatusStr = "confirmed"
	Prepare OrderStatusStr = "prepared"
	Deliver OrderStatusStr = "delivered"
)

func changeOrderStatusStr(status OrderStatusStr) {
	fmt.Println("changing order status to", status)
}

func main() {

	changeOrderStatus(Received)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Prepared)
	changeOrderStatus(Delivered)

	changeOrderStatusStr(Receive)
	changeOrderStatusStr(Confirm)
	changeOrderStatusStr(Prepare)
	changeOrderStatusStr(Deliver)

}
