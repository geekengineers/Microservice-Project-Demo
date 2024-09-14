package article_service

import (
	"context"
	"time"

	dto_object "github.com/tahadostifam/go-dto-object"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/article"
	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/dto"
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
	err := dto_object.Validate(dto.CreateArticleDto{Title: title, Description: description, Content: content})
	if err != nil {
		return nil, err
	}

	articleModel := &article.Article{
		Title:       title,
		Description: description,
		Content:     content,
		CoverImage:  "",
	}

	articleModel, err = s.requirements.Repo.Create(ctx, articleModel)
	if err != nil {
		return nil, ErrCreation
	}

	return articleModel, nil
}

// Delete implements Api.
func (s *Service) Delete(ctx context.Context, id int64) (*article.Article, error) {
	panic("unimplemented")
}

// Get implements Api.
func (s *Service) Get(ctx context.Context, id int64) (*article.Article, error) {
	panic("unimplemented")
}

// Search implements Api.
func (s *Service) Search(ctx context.Context, title string) ([]article.Article, error) {
	panic("unimplemented")
}

// Update implements Api.
func (s *Service) Update(ctx context.Context, id int64, changes *article.Article) (*article.Article, error) {
	panic("unimplemented")
}
