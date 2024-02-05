package models

type Order struct{
	Id uint `json:"id" gorm:"PrimaryKey"`
	ProductId []Product `json:"productId" gorm:"foreignKey:"Id"`
	UserId User `json:"userId" gorm:foreignKey:"Id"`

}