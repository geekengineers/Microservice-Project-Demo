package grpc_transformer

import (
	"time"

	article_domain "github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/article"
	"github.com/tahadostifam/go-hexagonal-architecture/protobuf/article"
)

func GrpcArticleToDomain(ar *article.Article) *article_domain.Article {
	return &article_domain.Article{
		ID:          ar.Id,
		Title:       ar.Title,
		Description: ar.Description,
		Content:     ar.Content,
		CoverImage:  ar.CoverImage,
		PublishedAt: time.Unix(ar.PublishAt, 0),
	}
}

func DomainToGrpcArticle(ar *article_domain.Article) *article.Article {
	return &article.Article{
		Id:          ar.ID,
		Title:       ar.Title,
		Description: ar.Description,
		Content:     ar.Content,
		CoverImage:  ar.CoverImage,
		PublishAt:   ar.PublishedAt.Unix(),
	}
}

func DomainToGrpcArticles(ar []article_domain.Article) []*article.Article {
	var tr []*article.Article

	for _, v := range ar {
		tr = append(tr, DomainToGrpcArticle(&v))
	}

	return tr
}