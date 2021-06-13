package main

import (
	"encoding/json"
	"fmt"
	"grapql-sample/model"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	tweet := model.Tweet{}
	tweets := tweet.Populate() 

	// Create Graphql Field Type

	var authorType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Author",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"twitterID": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var tweetType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Tweet",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"author": &graphql.Field{
					Type: authorType,
				},
				"dateTime": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
				"likes": &graphql.Field{
					Type: graphql.String,
				},
				"retweets": &graphql.Field{
					Type: graphql.String,
				},
				"comments": &graphql.Field{
					Type: graphql.NewList(graphql.String),
				},
			},
		},
	)
	//Define Schema
	fields := graphql.Fields{
		"tweet": &graphql.Field{
			Type:        tweetType,
			Description: "Get a tweet by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					for _, tweet := range tweets {
						if tweet.ID == id {
							return tweet, nil
						}
					}
				}
				return nil, nil
			},
		},
	}

	rq := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rq)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new GraphQL schema, err %v ", err)

	}

	query := `
	{
		tweet(id:1){
			author{
				name
				twitterID
			}
			dateTime
			content
			likes
			retweets 
		}
	}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if r.Errors != nil {

		log.Fatalf("Failed to execute graphql operation , errors : %+v", r.Errors)
	} 
	js, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("Failed to JSON Marshal : err : %v", err)
	}
	fmt.Printf("The json val is %s\n",js)
}
