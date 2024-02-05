package database

import (
	"log"
	"secweb/practice"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	
)

type SecendInstance struct{
	SecDb *gorm.DB
}
var Secenddatabase SecendInstance
func MakeSecendDb(){
	dsn:="root:sourav@90###@tcp(localhost:3306)/atma?parseTime"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Println("db connnection failed check secend db")
		
	}
	Secenddatabase.SecDb=db
	if err:=db.AutoMigrate(&practice.Mytrainer{},&practice.MyPokemon{});err!=nil{
		log.Println("problem in creating table in secend db")
		log.Fatal(err)
	}
	
	

	
}