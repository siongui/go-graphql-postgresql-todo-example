package graph

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

// A private key for context that only this package can access.
// This is important to prevent collisions between different context uses
var userCtxKey = &contextKey{name: "user"}

type contextKey struct {
	name string
}

func GetCtxKey() interface{} {
	return userCtxKey
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(userCtxKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func logAuthorizationHeader(ctx context.Context, logger log.Logger) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		logger.Log("logAuthorizationHeader_err", err)
	}

	if len(gc.Request.Header["Authorization"]) > 0 {
		logger.Log("Authorization", gc.Request.Header["Authorization"][0])
	}
}
