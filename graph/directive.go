package graph

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

func LogAuthorizationHeader(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return
	}

	if len(gc.Request.Header["Authorization"]) > 0 {
		log.Printf("Authorization: %s\n", gc.Request.Header["Authorization"][0])
	}

	return next(ctx)
}
