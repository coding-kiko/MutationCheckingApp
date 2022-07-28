package main

import (
	"os"

	"github.com/coding-kiko/MutantCheckingApp/pkg/mutant"
)

var (
	// Api
	ApiPort = os.Getenv("API_PORT")
	// DataBase
	DbUser = os.Getenv("POSTGRES_USER")
	DbPwd  = os.Getenv("POSTGRES_PWD")
	DbHost = os.Getenv("POSTGRES_HOST")
	DbPort = os.Getenv("POSTGRES_PORT")
	DbName = os.Getenv("POSTGRES_DB")
)

func main() {
	app := mutant.App{
		ApiPort: ApiPort,
		DbUser:  DbUser,
		DbPwd:   DbPwd,
		DbPort:  DbPort,
		DbHost:  DbHost,
		DbName:  DbName,
	}

	app.Init()
}
