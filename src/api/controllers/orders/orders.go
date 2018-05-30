package orders

import (
    "github.com/gin-gonic/gin"
    "strconv"
    "github.com/emikohmann/accelerator-api/src/api/services/orders"
    "net/http"
)

// localhost:8080/orders/34567

func GetOrder(c *gin.Context) {

    orderId, err := strconv.ParseInt(c.Param("orderId"), 10, 64)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        return
    }

    orderInformation, err := orders.GetOrderInformation(orderId)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
    }

    c.JSON(http.StatusOK, orderInformation)
}