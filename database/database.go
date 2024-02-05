package database

import (
	
	"log"
	"secweb/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type DbInstance struct{
	Db *gorm.DB
}
var Database DbInstance
 
func ConnectDb(){
	dsn:="root:sourav@90###@tcp(localhost:3306)/MyGoDb?parseTime=true"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal("connecting failed to db")

	}
	Database=DbInstance{Db:db}
	if err=db.AutoMigrate(&models.User{},&models.Product{});err!=nil{
		log.Println("table not created")
	}
	
	


}