package types

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userId": &graphql.Field{
			Type: graphql.String,
		},
		"userName": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"imageUrl": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "user",
	Fields: graphql.InputObjectConfigFieldMap{
		"userId": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"userName": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"imageUrl": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
	Description: "user input type",
})