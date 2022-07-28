package mutant

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type App struct {
	ApiPort string
	DbUser  string
	DbPwd   string
	DbPort  string
	DbHost  string
	DbName  string
}

func (a *App) Init() {
	postgresConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s%s", a.DbUser, a.DbPwd, a.DbHost, a.DbPort, a.DbName, "?sslmode=disable")
	postgresDb, err := sql.Open("postgres", postgresConnString)
	if err != nil {
		log.Fatalln(err)
	}
	defer postgresDb.Close()

	lab := NewLaboratory(postgresDb)
	service := NewMutantService(lab)
	handlers := NewHandlers(service)

	mux := RouteHandlers(handlers)
	fmt.Println("listening")
	log.Fatalln(http.ListenAndServe("0.0.0.0:"+a.ApiPort, mux))
}

func RouteHandlers(handlers Handlers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/mutant", handlers.CheckMutationHandler)
	mux.HandleFunc("/stats", handlers.StatsHandler)

	return mux
}
