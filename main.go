package main

import (
	"log"

	"secweb/database"
	"secweb/service"

	// "github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func main(){
	database.ConnectDb()
	database.MakeSecendDb()
	app :=*fiber.New()
	app.Use(cors.New())
	
	callRotes(&app)
	defer func(){
		sqlDb,_:=database.Database.Db.DB()
		err:=sqlDb.Close()
		if err!=nil{
			log.Println("db not close")
		}

	}()
	defer func(){
	    Ndb,_:=database.Secenddatabase.SecDb.DB()
		err:=Ndb.Close()
		if err!=nil{
			log.Println("not work secend db")
		}
	}()
	
	log.Fatal(app.Listen(":3000"))

}
func callRotes(app *fiber.App){
	app.Post("/users",service.CreateUserWithProduct)
	app.Post("/trainer",service.Combine)
	app.Get("/getAll",service.GetAll)
	app.Get("/getId/:id",service.FindById)
	app.Get("/getName/:name",service.FindByName)
	app.Post("/updateTrainer",service.UpdateTrainerById)
	app.Post("/updatePokemon",service.UpdatePokemonById)
	app.Get("/home",func(c *fiber.Ctx)error{
		return c.Render("forntend/index.html",fiber.Map{"title":"hello",})
	})
	app.Get("/load",func(c *fiber.Ctx)error{
		return c.Render("forntend/secend.html",fiber.Map{})
	})



}
