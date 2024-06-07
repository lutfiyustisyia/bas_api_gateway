package model

import "time"

type Transaction struct {
	Id              string `gorm:"primaryKey"`
	Account_id      string `gorm:"foreignkey"`
	Bank_id         string `gorm:"foreignkey"`
	Amount          int    `gorm:"column:amount"`
	TransactionDate *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
