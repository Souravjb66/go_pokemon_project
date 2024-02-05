package models


type User struct{
	Id uint `json:"id" gorm:"primary_Key"`
	Email string `json:"email" gorm:"unique;not null"`
	MyProduct []Product `json:"myproduct" gorm:"foreignKey:UserId"` // this is foreign key of product with from user one to many relation , so foreign key use in product table and userid is a field in product struct where user id stored
}