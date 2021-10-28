package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/machinebox/graphql"
)

func createProgram(w http.ResponseWriter, r *http.Request) {
	var newProgram program

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(reqBody, &newProgram)

	newProgram.ID = uuid.NewString()

	if newProgram.Name == "" {
		newProgram.Name = "No_name" + newProgram.ID
	}

	graphqlRequest := graphql.NewRequest(`
		mutation ($id: String!, $name: String!, $code: String!) {
			addProgram(input: {id: $id, name: $name, code: $code}) {
				program {
					id
					name
				}
			}
		}
	`)

	graphqlRequest.Var("id", newProgram.ID)
	graphqlRequest.Var("name", newProgram.Name)
	graphqlRequest.Var("code", newProgram.Code)
	var graphqlResponse interface{}

	client.Run(context.Background(), graphqlRequest, &graphqlResponse)

	render.JSON(w, 201, graphqlResponse)
}

func getPrograms(w http.ResponseWriter, r *http.Request) {
	graphqlRequest := graphql.NewRequest(`
		{
			queryProgram {
				id
				name
				code
			}
		}
	`)

	var graphqlResponse interface{}
	err := client.Run(context.Background(), graphqlRequest, &graphqlResponse)

	if err != nil {
		panic(err)
	}

	render.JSON(w, 200, graphqlResponse)
}

func getProgramById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	graphqlRequest := graphql.NewRequest(`
		query MyQuery ($id: String!){
			queryProgram(filter: {id: {eq: $id}}) {
				id
				name
				code
			}
		}
	`)

	graphqlRequest.Var("id", id)
	var graphqlResponse interface{}

	client.Run(context.Background(), graphqlRequest, &graphqlResponse)

	render.JSON(w, 200, graphqlResponse)
}
