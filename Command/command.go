package Command

import (
	"fmt"
	"strconv"
)

type Order interface {
	execute()
}

type Stock struct {
	name     string
	quantity int
}

func newStock(name string, quantity int) *Stock {
	return &Stock{name, quantity}
}

func (s Stock) Buy() string {
	return "Buy: " + s.name + ":" + strconv.Itoa(s.quantity)
}

func (s Stock) Sell() string {
	return "Sell: " + s.name + ":" + strconv.Itoa(s.quantity)
}

type BuyStock struct {
	Stock
}

func newBuyStock(s *Stock) *BuyStock {
	if s == nil {
		return &BuyStock{}
	}
	return &BuyStock{*s}
}

func (b *BuyStock) execute() {
	fmt.Println(b.Buy())
}

type SellStock struct {
	Stock
}

func newSellStock(s *Stock) *SellStock {
	if s == nil {
		return &SellStock{}
	}
	return &SellStock{*s}
}

func (s *SellStock) execute() {
	fmt.Println(s.Sell())
}

type Broker []Order

func NewBroker(o ...Order) *Broker {
	var b Broker
	b = append(b, o...)
	return &b
}

func (b *Broker) TakeOrder(o ...Order) {
	if b == nil {
		*b = make([]Order, 0)
	}
	*b = append(*b, o...)
}

func (b Broker) PlaceOrders() {
	for _, order := range b {
		order.execute()
	}
}
