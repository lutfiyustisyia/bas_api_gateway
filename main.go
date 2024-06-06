package main

import (
	"github.com/gin-gonic/gin"

	"api_gateway/handler"
)

// main  excercise 5 juni 2024
func main() {
	r := gin.Default()

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/remove/:id", handler.NewAccount().RemoveAccount)
	accountRoute.POST("/getbalance", handler.NewAccount().GetBalance)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transferbank", handler.NewTransaction().TransferBank)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
