package grpc_adapter

import (
	"context"
	"fmt"
	"net"

	"github.com/geekengineers/Microservice-Project-Demo/protobuf/article"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/interceptor"
	grpc_transformer "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/adapters/primary/grpc/transformer"
	article_domain "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/domain/article"
	article_service "github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/services/article"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type App struct {
	articleService article_service.Api
	host           string
	port           int
	server         *grpc.Server
}

type articleServerImpl struct {
	article.UnimplementedArticleServiceServer
	articleService article_service.Api
}

func (a articleServerImpl) Create(ctx context.Context, req *article.CreateRequest) (*article.CreateResponse, error) {
	ar, err := a.articleService.Create(ctx, req.Title, req.Description, req.Content, req.CoverImage)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &article.CreateResponse{
		Article: grpc_transformer.DomainToGrpcArticle(ar),
	}

	return res, nil
}

func (a articleServerImpl) Delete(ctx context.Context, req *article.DeleteRequest) (*article.DeleteResponse, error) {
	err := a.articleService.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res := &article.DeleteResponse{}

	return res, nil
}

func (a articleServerImpl) Find(ctx context.Context, req *article.FindRequest) (*article.FindResponse, error) {
	ar, err := a.articleService.Find(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &article.FindResponse{
		Article: grpc_transformer.DomainToGrpcArticle(ar),
	}

	return res, nil
}

func (a articleServerImpl) Search(ctx context.Context, req *article.SearchRequest) (*article.SearchResponse, error) {
	articles, err := a.articleService.Search(ctx, req.Input)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &article.SearchResponse{
		Articles: grpc_transformer.DomainToGrpcArticles(articles),
	}

	return res, nil
}

func (a articleServerImpl) Update(ctx context.Context, req *article.UpdateRequest) (*article.UpdateResponse, error) {
	changes := article_domain.Article{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		CoverImage:  req.CoverImage,
	}

	ar, err := a.articleService.Update(ctx, req.Id, &changes)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	res := &article.UpdateResponse{
		Article: grpc_transformer.DomainToGrpcArticle(ar),
	}

	return res, nil
}

func NewGrpcServer(articleService article_service.Api, authServiceUrl string, host string, port int) *App {
	authInterceptor := interceptor.AuthInterceptor(authServiceUrl)
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(authInterceptor))

	article.RegisterArticleServiceServer(s, articleServerImpl{articleService: articleService})

	return &App{articleService, host, port, s}
}

func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.host, a.port))
	if err != nil {
		return err
	}

	err = a.server.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
