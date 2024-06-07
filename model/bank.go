package model

type Bank struct {
	Bank_code string `gorm:"primaryKey"`
	Name      string
	Address   string
}

func (a *Bank) TableName() string {
	return "bank"
}
