package model

type CartItem struct {
	CartItemID int64
	Book       *Book
	Count      int64
	Amount     float64
	CartID     string
}

func (this *CartItem) GetAmount() float64 {
	price := this.Book.Price
	return float64(this.Count) * price
}
