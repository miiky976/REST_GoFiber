package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	user     = "miiky"
	pass     = "datata"
	serv     = "127.0.0.1:3306"
	database = "Tareas"
	DSN      = user + ":" + pass + "@tcp(" + serv + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB       *gorm.DB
	err      error
)

func DBConection() {
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Conexion existosa")
	}
}
