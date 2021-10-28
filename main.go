package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/machinebox/graphql"

	renderPkg "github.com/unrolled/render"
)

var render *renderPkg.Render
var client *graphql.Client

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Golang programs API")
}

func ProgramsHandler(router chi.Router) {
	router.Get("/", getPrograms)
	router.Post("/", createProgram)
	router.Get("/{id}", getProgramById)
}

func main() {
	render = renderPkg.New()
	router := chi.NewRouter()
	client = graphql.NewClient("https://blue-surf-420063.us-east-1.aws.cloud.dgraph.io/graphql")
	router.Get("/", HomeHandler)
	router.Route("/programs", ProgramsHandler)
	http.Handle("/", router)

	http.ListenAndServe(":3000", router)
}
