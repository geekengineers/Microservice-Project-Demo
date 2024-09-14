package article_service

import (
	"context"
	"time"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/article"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/ports"
)

const AccessTokenTTL = 20 * 24 * time.Hour

type Requirements struct {
	Repo ports.ArticleRepositorySecondaryPort
}

type Service struct {
	requirements *Requirements
}

func NewService(requirements *Requirements) *Service {
	return &Service{requirements}
}

func (s *Service) Create(ctx context.Context, title, description, content string) (*article.Article, error) {
	articleModel := &article.Article{
		Title:       title,
		Description: description,
		Content:     content,
		CoverImage:  "",
	}
	articleModel, err := s.requirements.Repo.Create(ctx, articleModel)
	if err != nil {
		return nil, ErrCreation
	}

	return articleModel, nil
}
