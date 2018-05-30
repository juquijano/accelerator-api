package addresses

type Address struct {
    AddressId int64  `json:"address_id"`
    City      string `json:"city"`
    Street    string `json:"street"`
    Number    int    `json:"number"`
}
