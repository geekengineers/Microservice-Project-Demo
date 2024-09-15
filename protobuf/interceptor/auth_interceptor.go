package interceptor

import (
	"context"
	"sync"
	"time"

	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var authServiceUrl string

type authKey struct{}

var AuthKey = &authKey{}

type GrpcInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error)

func AuthInterceptor(url string) GrpcInterceptor {
	authServiceUrl = url

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok || md.Len() == 0 || len(md["X-Access-Token"]) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing access token header")
		}

		accessToken := md.Get("X-Access-Token")[0]

		client := authServicePool.Get().(auth.AuthClient)

		res, err := client.Authenticate(ctx, &auth.AuthenticateRequest{
			AccessToken: accessToken,
		})
		if err != nil && res.User != nil {
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		}

		newCtx := context.WithValue(ctx, AuthKey, res.User.Id)

		return handler(newCtx, req)
	}
}

func establishAuthServiceConn() (*grpc.ClientConn, error) {
	return grpc.NewClient(authServiceUrl, grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
}

var authServicePool = &sync.Pool{
	New: func() any {
		var conn *grpc.ClientConn
		var err error

		for {
			conn, err = establishAuthServiceConn()
			if err != nil {
				time.Sleep(250 * time.Millisecond)
				continue
			}

			break
		}

		client := auth.NewAuthClient(conn)

		return client
	},
}
