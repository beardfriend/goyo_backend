package libs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ENV Env

type Env struct {
	Port string
	Jwt  string
	Database
	API
	Redis
}

type Database struct {
	Host string
	Pass string
	User string
	Port string
	DB   string
}

type API struct {
	ThisProject string
}

type Redis struct {
	Host string
	Pass string
	Port string
}

func CheckEnv() (string, error) {
	s := os.Getenv("GO_ENV")
	url := "./.env.dev"
	mode := "development"

	if s == "production" {
		url = "./.env.prod"
		mode = "production"
	}

	if err := godotenv.Load(url); err != nil {
		return "", fmt.Errorf("%s", "Please Setting Env File")
	}
	return mode, nil
}

func ReadEnv(mode string) {
	var result map[string]string
	var error error
	if mode == "production" {
		str, err := godotenv.Read("./.env.prod")
		result = str
		error = err
	}
	if mode == "test" {
		str, err := godotenv.Read("../../.env.dev")
		result = str
		error = err
	} else {
		str, err := godotenv.Read("./.env.dev")
		result = str
		error = err
	}
	if error != nil {
		panic(error)
	}

	env_database := &Database{
		Host: result["DATABASE_HOST"],
		Pass: result["DATABASE_PASS"],
		User: result["DATABASE_USER"],
		Port: result["DATABASE_PORT"],
		DB:   result["DATABASE_DB"],
	}

	env_api := &API{
		ThisProject: result["API_THISPROJECT"],
	}

	env_redis := &Redis{
		Host: result["REDIS_HOST"],
		Pass: result["REDIS_PASS"],
		Port: result["REDIS_PORT"],
	}

	env := &Env{
		Port:     result["PORT"],
		Jwt:      result["JWT"],
		Database: *env_database,
		API:      *env_api,
		Redis:    *env_redis,
	}

	ENV = *env
}
