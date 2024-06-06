package model

type Account struct {
	Account_ID string `gorm:"primaryKey"`
	Username   string
	Password   string
	Name       string
}

func (a *Account) TableName() string {
	return "account"
}
