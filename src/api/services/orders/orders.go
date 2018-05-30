package orders

import (
    ordersDomain "github.com/emikohmann/accelerator-api/src/api/domain/orders"
    itemsDomain "github.com/emikohmann/accelerator-api/src/api/domain/items"
    addressesDomain "github.com/emikohmann/accelerator-api/src/api/domain/addresses"
    paymentsDomain "github.com/emikohmann/accelerator-api/src/api/domain/payments"
    "fmt"
)

func GetOrderInformation(orderId int64) (*ordersDomain.OrderInformation, error) {
    orderData, err := getOrderData(orderId)
    if err != nil {
        return nil, err
    }

    return processOrderData(orderData)
}

func getOrderData(orderId int64) (*ordersDomain.OrderData, error) {
    return &ordersDomain.OrderData{
        OrderId:   orderId,
        ItemId:    1234,
        AddressId: 1234,
        PaymentId: 1234,
    }, nil
}

func processOrderData(orderData *ordersDomain.OrderData) (*ordersDomain.OrderInformation, error) {

    item, err := getItem(orderData.ItemId)
    if err != nil {
        return nil, err
    }

    address, err := getAddress(orderData.AddressId)
    if err != nil {
        return nil, err
    }

    payment, err := getPayment(orderData.PaymentId)
    if err != nil {
        return nil, err
    }

    return &ordersDomain.OrderInformation{
        OrderID: orderData.OrderId,
        ItemName: item.Name,
        ItemPrice: fmt.Sprintf("$%2f", item.Price),
        Address: fmt.Sprintf("%s %d - %s", address.Street, address.Number, address.City),
        Payment: fmt.Sprintf("$%2f with %s", payment.Amount, payment.PaymentMethod),
    }, nil
}

func getItem(itemId int64) (*itemsDomain.Item, error) {
    return &itemsDomain.Item{
        ItemId:   itemId,
        Name:     "Mouse",
        Price:    500,
        Quantity: 10,
    }, nil
}

func getAddress(addressId int64) (*addressesDomain.Address, error) {
    return &addressesDomain.Address{
        AddressId: addressId,
        City:      "Cordoba",
        Street:    "Av. La Voz del Interior",
        Number:    7000,
    }, nil
}

func getPayment(paymentId int64) (*paymentsDomain.Payment, error) {
    return &paymentsDomain.Payment{
        PaymentId:     paymentId,
        PaymentMethod: "Credit Card",
        Amount:        500,
    }, nil
}
