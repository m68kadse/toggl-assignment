package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/m68kadse/toggl-assignment/dao/sqlite"
	"github.com/m68kadse/toggl-assignment/graph"
)

const defaultPort = "3000"
const defaultFile = "questions.db"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	file := os.Getenv("FILE")
	if file == "" {
		file = defaultFile
	}

	//create SQLiteDAO
	dao, err := sqlite.NewDAO(file)
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		QuestionDAO: dao}}))

	http.Handle("/", srv)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
