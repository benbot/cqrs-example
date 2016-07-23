package query

import (
	"cqrs-example/player"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"player": &graphql.Field{
				Type: player.PlayerType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//TODO make some store to read from
				},
			},
		},
	},
)

func execQuery(query string, schema graphql.Schema) *graphql.Result {

}
