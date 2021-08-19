package Command

import "testing"

func TestCommandStockOrder(t *testing.T) {
	s := newStock("stock", 10)
	buyStock := newBuyStock(s)
	sellStock := newSellStock(s)

	broker := NewBroker(buyStock, sellStock)

	broker.PlaceOrders()

	s2 := newStock("stock2", 20)
	buyStock2 := newBuyStock(s2)
	sellStock2 := newSellStock(s2)
	broker.TakeOrder(buyStock2, sellStock2)
	broker.PlaceOrders()
}
