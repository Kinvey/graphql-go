package graphql

import (
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/internal/query"
	"github.com/graph-gophers/graphql-go/types"
)

func ParseQuery(queryString string) (*types.ExecutableDefinition, *errors.QueryError) {
	return query.Parse(queryString)
}
