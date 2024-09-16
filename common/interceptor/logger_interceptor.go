package interceptor

import (
	"context"
	"log"

	"connectrpc.com/connect"
)

func LoggerInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			log.Printf("[GRPC REQUEST][%s] %s\n", req.HTTPMethod(), req.Spec().Schema)

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}
