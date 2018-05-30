package items

type Item struct {
    ItemId   int64   `json:"item_id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}
