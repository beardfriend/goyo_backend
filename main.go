package main

import (
	"log"

	"goyo/libs"
	"goyo/server"
)

func main() {
	mode, err := libs.CheckEnv()
	if err != nil {
		log.Fatalln(err)
	}
	libs.ReadEnv(mode)

	server.InitGin()
}
