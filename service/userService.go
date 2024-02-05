package service

import (
	// "secweb/database"
	"log"
	"secweb/database"
	"secweb/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUserWithProduct(c *fiber.Ctx)error{
	
	type MyUser struct {
		Email    string `json:"email" gorm:"not null"`
		Products []models.Product `json:"products"`
	}

	var myUser MyUser
	if err:=c.BodyParser(&myUser);err!=nil{
		log.Println("problem in body parse")

	}
	
	ForUser:=models.User{
		Email: myUser.Email,
		MyProduct: myUser.Products,
	}
	log.Println(myUser.Products)
	if err:=database.Database.Db.Create(&ForUser);err!=nil{
		log.Println(err)
	}
	
	return c.Status(200).JSON(ForUser)
	
}