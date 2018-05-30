package app

import (
    "github.com/gin-gonic/gin"
    pingController "github.com/emikohmann/accelerator-api/src/api/controllers/ping"
    ordersController "github.com/emikohmann/accelerator-api/src/api/controllers/orders"
)

func StartApp() {

    router := gin.Default()

    router.GET("/ping", pingController.Ping)
    router.GET("/orders/:orderId", ordersController.GetOrder)

    router.Run(":8080")

}
