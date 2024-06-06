package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-micro/plugins/v4/client/grpc"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/client"

	"api_gateway/handler"
	"api_gateway/proto"
)

// main  excercise 5 juni 2024
func main() {
	r := gin.Default()
	addServiceTransactionOpt := client.WithAddress(":8080")
	clientSrvTransaction := grpc.NewClient()
	srvTransaction := micro.NewService(
		micro.Client(clientSrvTransaction),
	)
	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	grpc.NewClient()

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/remove/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/getbalance", handler.NewAccount().GetBalance)

	transactionRoute := r.Group("/transaction")
	transactionRoute.GET("/get", func(g *gin.Context) {
		ClientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).Call(context.Background(), &proto.CallRequest{
			Name: "Lupi",
		}, addServiceTransactionOpt)
		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		g.JSON(http.StatusOK, gin.H{
			"data": ClientResponse,
		})
	})
	transactionRoute.POST("/transferbank", handler.NewTransaction().TransferBank)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)
	r.Run(":9888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
