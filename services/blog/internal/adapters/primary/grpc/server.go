package grpc_adapter

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/geekengineers/Microservice-Project-Demo/common/interceptor"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/article"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/article/articleconnect"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/config"
	article_domain "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/domain/article"

	grpc_transformer "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/adapters/primary/grpc/transformer"
	article_service "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/services/article"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type App struct {
	articleService article_service.Api
	host           string
	port           int
	mux            *http.ServeMux
}

type articleServerImpl struct {
	articleService article_service.Api
}

func (a articleServerImpl) Create(ctx context.Context, req *connect.Request[article.CreateRequest]) (*connect.Response[article.CreateResponse], error) {
	ok := interceptor.RoleRequired(ctx, "admin")

	if ok {
		ar, err := a.articleService.Create(ctx, req.Msg.Title, req.Msg.Description, req.Msg.Content, req.Msg.CoverImage)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		res := &connect.Response[article.CreateResponse]{
			Msg: &article.CreateResponse{
				Article: grpc_transformer.DomainToGrpcArticle(ar),
			},
		}

		return res, nil
	}

	return nil, connect.NewError(connect.CodeUnauthenticated, interceptor.CodePermissionDenied)
}

func (a articleServerImpl) Delete(ctx context.Context, req *connect.Request[article.DeleteRequest]) (*connect.Response[article.DeleteResponse], error) {
	ok := interceptor.RoleRequired(ctx, "admin")

	if ok {
		err := a.articleService.Delete(ctx, req.Msg.Id)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}

		res := &connect.Response[article.DeleteResponse]{
			Msg: &article.DeleteResponse{},
		}

		return res, nil
	}

	return nil, connect.NewError(connect.CodeUnauthenticated, interceptor.CodePermissionDenied)
}

func (a articleServerImpl) Find(ctx context.Context, req *connect.Request[article.FindRequest]) (*connect.Response[article.FindResponse], error) {
	ar, err := a.articleService.Find(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := &connect.Response[article.FindResponse]{
		Msg: &article.FindResponse{
			Article: grpc_transformer.DomainToGrpcArticle(ar),
		},
	}

	return res, nil
}

func (a articleServerImpl) Search(ctx context.Context, req *connect.Request[article.SearchRequest]) (*connect.Response[article.SearchResponse], error) {
	articles, err := a.articleService.Search(ctx, req.Msg.Input)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	res := &connect.Response[article.SearchResponse]{
		Msg: &article.SearchResponse{
			Articles: grpc_transformer.DomainToGrpcArticles(articles),
		},
	}

	return res, nil
}

func (a articleServerImpl) Update(ctx context.Context, req *connect.Request[article.UpdateRequest]) (*connect.Response[article.UpdateResponse], error) {
	ok := interceptor.RoleRequired(ctx, "admin")

	if ok {
		changes := article_domain.Article{
			Title:       req.Msg.Title,
			Description: req.Msg.Description,
			Content:     req.Msg.Content,
			CoverImage:  req.Msg.CoverImage,
		}

		ar, err := a.articleService.Update(ctx, req.Msg.Id, &changes)
		if err != nil {
			return nil, connect.NewError(connect.CodeInvalidArgument, err)
		}

		res := &connect.Response[article.UpdateResponse]{
			Msg: &article.UpdateResponse{
				Article: grpc_transformer.DomainToGrpcArticle(ar),
			},
		}

		return res, nil
	}

	return nil, connect.NewError(connect.CodeUnauthenticated, interceptor.CodePermissionDenied)
}

func NewGrpcServer(articleService article_service.Api, authServiceUrl string, host string, port int) *App {
	loggerInterceptor := interceptor.LoggerInterceptor()
	authInterceptor := interceptor.NewAuthInterceptor(authServiceUrl)
	interceptorList := []connect.Interceptor{loggerInterceptor}
	if config.CurrentEnv != config.Test {
		interceptorList = append(interceptorList, authInterceptor)
	}
	interceptors := connect.WithInterceptors(interceptorList...)

	mux := http.NewServeMux()

	path, handler := articleconnect.NewArticleServiceHandler(articleServerImpl{articleService}, interceptors)
	mux.Handle(path, handler)

	return &App{articleService, host, port, mux}
}

func (a *App) Run() error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", a.host, a.port), h2c.NewHandler(a.mux, &http2.Server{}))
	if err != nil {
		return err
	}

	return nil
}
