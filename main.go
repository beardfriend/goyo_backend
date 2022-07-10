package main

import (
	"log"

	"goyo/libs"
	"goyo/models"
	"goyo/models/users"
	"goyo/server/echo"
	"goyo/server/mariadb"
)

func main() {
	mode, err := libs.CheckEnv()
	if err != nil {
		log.Fatalln(err)
	}
	libs.ReadEnv(mode)
	db := mariadb.GetInstance()
	db.AutoMigrate(&models.Health{}, &users.User{}, &users.Google{}, &users.Kakao{}, &users.Normal{}, &users.Social{})
	echo.InitEcho()
}
