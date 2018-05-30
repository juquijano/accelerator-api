package orders

type OrderData struct {
    OrderId   int64
    ItemId    int64
    AddressId int64
    PaymentId int64
}

type OrderInformation struct {
    OrderID   int64  `json:"order_id"`
    ItemName  string `json:"item_name"`
    ItemPrice string `json:"item_price"`
    Address   string `json:"address"`
    Payment   string `json:"payment"`
}
