package service

import (
	"log"
	"secweb/database"
	"secweb/practice"

	"github.com/gofiber/fiber/v2"
)

func SendToPoke(c *fiber.Ctx)error{
	type Mypoke struct{
		Name string `json:"name" gorm:"not null"`
		Special string `json:"special" gorm:"not null"`
		Pokemon practice.MyPokemon `json:"pokemon" gorm:"not null"`
	}
	// var myPoke Mypoke 
	myPoke:=Mypoke{}  
	err:=c.BodyParser(&myPoke)
	if err!=nil{
		log.Println("practice body parse have issue")
		log.Println(err)
		log.Println(myPoke)
	}
	Fortrainer:=practice.Mytrainer{
		Name:myPoke.Name,
		Special: myPoke.Special,
		// Pokemon: myPoke.Pokemon,
	}
	database.Secenddatabase.SecDb.Create(&Fortrainer)
	return c.Status(200).JSON(Fortrainer)


}
func Combine(c *fiber.Ctx)error{
	type MyComb struct{
		Name string `json:"name" gorm:"not null"`
		Special string `json:"special" gorm:"not null"`
		Pokemons []practice.MyPokemon `json:"pokemons" gorm:"not null"`
	}
	myComb:=MyComb{}
	if err:=c.BodyParser(&myComb);err!=nil{
		log.Println(err)
	}
	ForTrainerPokemon:=practice.Mytrainer{
		Name:myComb.Name,
		Special:myComb.Special,
		Pokemon: myComb.Pokemons,
	}
	log.Println(ForTrainerPokemon)
	database.Secenddatabase.SecDb.Create(&ForTrainerPokemon)
	return c.Status(200).JSON(ForTrainerPokemon)
}
func GetAll(c *fiber.Ctx)error{
	var trainers []practice.Mytrainer
	err:=database.Secenddatabase.SecDb.Preload("Pokemon").Find(&trainers)
	if err!=nil{
		log.Println("problem in getAll",err)
	}
	return c.Status(200).JSON(trainers)
}

func FindByName(c *fiber.Ctx)error{
	trainers:=practice.Mytrainer{}
	value:=c.Params("name")
	err:=database.Secenddatabase.SecDb.Preload("Pokemon").Find(&trainers,"name=?",value)
	if err!=nil{
		log.Println(err)
	}
	return c.Status(200).JSON(trainers)
}
func FindById(c *fiber.Ctx)error{
	trainers:=practice.Mytrainer{}
	valueId:=c.Params("id")
	err:=database.Secenddatabase.SecDb.Preload("Pokemon").Find(&trainers,"id=?",valueId)  //id use because its define in json
	if err!=nil{
		log.Println("find by id preload",err)
	}
	return c.Status(200).JSON(trainers)

}
func UpdateTrainerById(c *fiber.Ctx)error{
	myTrainer:=practice.Mytrainer{}
	type Trainer struct{
		Id uint `json:"id"`
		Name string `json:"name"`
		Special string `json:"special"`

	}
	trainer:=Trainer{}
	err:=c.BodyParser(&trainer)
	if err!=nil{
		log.Println("log in update trainer",err)
	}
	if trainer.Id==0{
		return c.JSON("id must be not 0")
	}
	valueId:=trainer.Id
	getTrainer,err:=findTrainer(&myTrainer,valueId)
	if err!=nil{
		log.Println("error in find function",err)
	}
	
	UpdateValue:=practice.Mytrainer{
		Id:getTrainer.Id,
		Name:trainer.Name,
		Special:trainer.Special,
	}
	log.Println(UpdateValue)
	log.Println(&UpdateValue)
	database.Secenddatabase.SecDb.Save(&UpdateValue)
	return c.JSON(UpdateValue)
}
func findTrainer(trainer *practice.Mytrainer,value uint)(*practice.Mytrainer,error){
	database.Secenddatabase.SecDb.Find(trainer,"id=?",value)
	log.Println(trainer)
	return trainer,nil
}
func UpdatePokemonById(c *fiber.Ctx)error{
	type Pokemon struct{
		Id uint
		Name string
	}
	myPokemon:=Pokemon{}
	err:=c.BodyParser(&myPokemon)
	if err!=nil{
		log.Println(err)
	}
	if myPokemon.Id==0{
		return c.JSON("give greter no")
	}
	getPokemon:=findPokemon(myPokemon.Id)
	savePokemon:=practice.MyPokemon{
		Id:getPokemon.Id,
		Name: myPokemon.Name,
	}
	database.Secenddatabase.SecDb.Save(&savePokemon)
	return c.Status(200).JSON(savePokemon)

}
func findPokemon(id uint)*practice.MyPokemon{
	pokemon:=practice.MyPokemon{}
	database.Secenddatabase.SecDb.Find(&pokemon,"id=?",id)
	return &pokemon
}
