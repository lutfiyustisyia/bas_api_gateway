package model

import "time"

type Transaction struct {
	Id               int     `gorm:"primaryKey`
	Account_id       string  `gorm:"foreignkey" json: account_id`
	Bank_id          string  `gorm:"foreignkey" json: bank_id`
	Amount           float64 `gorm:"column:amount" json: amount`
	Transaction_date *time.Time
}

func (a *Transaction) TableName() string {
	return "transaction"
}
