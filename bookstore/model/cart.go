package model

type Cart struct {
	CartID      string
	CartItems   []*CartItem
	TotalCount  int64
	TotalAmount float64
	UserID      int
}

func (this *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range this.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

func (this *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range this.CartItems {
		totalAmount = totalAmount + v.Amount
	}
	return totalAmount
}
