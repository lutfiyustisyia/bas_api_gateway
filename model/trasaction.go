package model

import "time"

type Transaction struct {
	Id              string `gorm:"primaryKey"`
	account_id      string `gorm:"foreignkey"`
	bank_id         string `gorm:"foreignkey"`
	Amount          int    `gorm:"column:amount"`
	TransactionDate *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
