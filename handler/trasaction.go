package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	TransferBank(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

type BodyPayloadTransaction struct{}

func (b *transactionImplement) TransferBank(g *gin.Context) {

	bodyPayloadTxn := model.Transaction{}
	err := g.BindJSON(&bodyPayloadTxn)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&bodyPayloadTxn)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Transaction Succesfull",
		"data":    bodyPayloadTxn,
	})
}

func GetAccounts(c *gin.Context) {
	orm := utils.NewDatabase().Orm
	db, err := orm.DB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
		return
	}
	defer db.Close()

	var accounts []model.Account
	if err := orm.Find(&accounts).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": accounts,
	})
}

func GetBanks(c *gin.Context) {
	orm := utils.NewDatabase().Orm
	db, err := orm.DB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to database",
		})
		return
	}
	defer db.Close()

	var banks []model.Bank
	if err := orm.Find(&banks).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": banks,
	})
}
