package main

import (
	"log"

	"goyo/libs"
	"goyo/server"
	"goyo/server/mariadb"
)

func main() {
	mode, err := libs.CheckEnv()
	if err != nil {
		log.Fatalln(err)
	}
	libs.ReadEnv(mode)
	mariadb.TableGenerate()
	server.InitGin()
}
