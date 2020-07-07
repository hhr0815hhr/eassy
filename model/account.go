package model

import "github.com/jinzhu/gorm"

type Account struct {
	*gorm.Model
	Phone string `gorm:"size:13;not null"`
	Pwd   string `gorm:"size:80;not null"`
}

var acc Account

func IsExist(accId uint) bool {
	dbRead.Where("id=?", accId).Find(&acc)
	return acc.ID != 0
}

func RegAccount(phone, pwd string) error {
	acc = Account{
		Phone: phone,
		Pwd:   pwd,
	}
	return dbWrite.Create(&acc).Error
}

func GetAccountByPhone(phone string) *Account {
	dbRead.Where("phone=?", phone).Find(&acc)
	return &acc
}
