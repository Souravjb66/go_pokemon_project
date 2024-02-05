package practice

type MyPokemon struct{
	Id uint `json:"id" gorm:"primary_Key"`
	Name string `json:"name" gorm:"not null"`
	Trainer []Mytrainer `json:"trainer" gorm:"many2many:poke_trainer"`
	
	
	
}