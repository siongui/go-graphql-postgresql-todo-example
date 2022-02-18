package graph

import (
	"context"

	"github.com/siongui/go-kit-gqlgen-postgres-todo-example/graph/generated"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-kit/log"
)

type directive struct {
	logger log.Logger
}

func (d *directive) logAuthorizationHeader(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return
	}

	if len(gc.Request.Header["Authorization"]) > 0 {
		d.logger.Log("Authorization", gc.Request.Header["Authorization"][0])
	}

	return next(ctx)
}

func (d *directive) logHeader(ctx context.Context, obj interface{}, next graphql.Resolver, header string) (res interface{}, err error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return
	}

	if len(gc.Request.Header[header]) > 0 {
		d.logger.Log(header, gc.Request.Header[header][0])
	}

	return next(ctx)
}

func NewDirectives(logger log.Logger) generated.DirectiveRoot {
	d := directive{logger: logger}

	return generated.DirectiveRoot{
		LogAuthorizationHeader: d.logAuthorizationHeader,
		LogHeader:              d.logHeader,
	}
}
