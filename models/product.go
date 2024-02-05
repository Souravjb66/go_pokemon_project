package models

type Product struct{
	Id uint `json:"id" gorm:"primary_Key"`
	ProductName string `json:"productname" gorm:"not null"`
	Price uint `json:"price" gorm:"not null"`
	UserId uint `json:"userid"`
}