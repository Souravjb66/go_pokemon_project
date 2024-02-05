package practice

type Mytrainer struct{
	Id uint `json:"id" gorm:"primary_Key"`
	Name string `json:"name" gorm:"not null"`
	Special string `json:"special" gorm:"not null"`
	Pokemon []MyPokemon `json:"pokemon" gorm:"many2many:poke_trainer"`
}