package data

type Order struct {
	Model
	UserID        uint           `schema:"user_id"`
	TotalMoney    float64        `schema:"total_money"`
	Status        string         `schema:"status"`
	PaymentID     uint           `schema:"payment_id"`
	Payment       *Payment       //has one payment method
	OrderProducts []OrderProduct //has many order products
	User          *User          //belong to user
}

type OrderProduct struct {
	Model
	OrderID   uint     `schema:"order_id"`
	ProductID uint     `schema:"product_id"`
	Quantity  uint     `schema:"quantity"`
	Product   *Product //belong to product
	Order     *Order   //belong to order
}
