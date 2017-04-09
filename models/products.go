package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Team struct {
	gorm.Model
	Name string
}

func GetAll() string {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=gouser dbname=godb sslmode=disable password=gopass")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Team{})

	var team Team
	db.First(&team, 1)
	return string(team.Name)
}

func Create() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=gouser dbname=godb sslmode=disable password=gopass")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Team{})

	db.Create(&Team{Name: "Phils"})
}
