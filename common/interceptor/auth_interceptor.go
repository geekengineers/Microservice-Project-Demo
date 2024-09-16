package interceptor

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"connectrpc.com/connect"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/auth/authconnect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var authServiceUrl string

type authKey struct{}

type AuthenticatedUser struct {
	UserID string
	Role   string
}

var AuthKey = &authKey{}

func Unauthenticated() error {
	return status.Errorf(codes.Unauthenticated, "unauthenticated")
}

func NewAuthInterceptor(url string) connect.UnaryInterceptorFunc {
	authServiceUrl = url

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			accessToken := req.Header().Get("X-Access-Token")
			if len(accessToken) != 0 {
				authService := authServicePool.Get().(authconnect.AuthServiceClient)

				res, err := authService.Authenticate(ctx, &connect.Request[auth.AuthenticateRequest]{
					Msg: &auth.AuthenticateRequest{
						AccessToken: accessToken,
					},
				})
				if err == nil {
					newCtx := context.WithValue(ctx, AuthKey, &AuthenticatedUser{
						UserID: res.Msg.User.Id,
						Role:   res.Msg.User.Role,
					})

					return next(newCtx, req)
				}

				return next(ctx, req)
			}

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

var authServicePool = &sync.Pool{
	New: func() any {
		client := authconnect.NewAuthServiceClient(http.DefaultClient, authServiceUrl, connect.WithClientOptions())

		return client
	},
}

func AuthRequired(ctx context.Context) bool {
	userId := ctx.Value(AuthKey)

	return userId == nil
}

func RoleRequired(ctx context.Context, expectedRole string) bool {
	userInfoAny := ctx.Value(AuthKey)

	if userInfoAny != nil {
		userInfo := userInfoAny.(*AuthenticatedUser)

		if userInfo.UserID == "" || userInfo.Role == "" {
			return false
		}

		if userInfo.Role != expectedRole {
			return false
		}
	}

	return false
}

var CodeAuthRequired = connect.NewError(connect.CodeInvalidArgument, errors.New("only authenticated users allowed"))
var CodePermissionDenied = connect.NewError(connect.CodeInvalidArgument, errors.New("your role is not allowed to do this action"))
