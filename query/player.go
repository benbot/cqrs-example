package query

import (
	"cqrs-example/player"

	"github.com/graphql-go/graphql"
)

var playerField = &graphql.Field{
	Type: player.PlayerType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
}
