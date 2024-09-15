package article_service

import (
	"context"

	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/domain/article"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/dto"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/ports"
	dto_object "github.com/tahadostifam/go-dto-object"
)

type Requirements struct {
	Repo ports.ArticleRepositorySecondaryPort
}

type Service struct {
	requirements *Requirements
}

func NewService(requirements *Requirements) *Service {
	return &Service{requirements}
}

func (s *Service) Create(ctx context.Context, title, description, content, coverImage string) (*article.Article, error) {
	err := dto_object.Validate(dto.CreateArticleDto{
		Title:       title,
		Description: description,
		Content:     content,
		CoverImage:  coverImage,
	})
	if err != nil {
		return nil, err
	}

	articleModel := &article.Article{
		Title:       title,
		Description: description,
		Content:     content,
		CoverImage:  coverImage,
	}

	articleModel, err = s.requirements.Repo.Create(ctx, articleModel)
	if err != nil {
		return nil, ErrCreation
	}

	return articleModel, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.requirements.Repo.Delete(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	return nil
}

func (s *Service) Find(ctx context.Context, id int64) (*article.Article, error) {
	article, err := s.requirements.Repo.Find(ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	return article, nil
}

func (s *Service) Search(ctx context.Context, title string) ([]article.Article, error) {
	articles, err := s.requirements.Repo.Search(ctx, title)
	if err != nil {
		return nil, ErrSearch
	}

	return articles, nil
}

func (s *Service) Update(ctx context.Context, id int64, changes *article.Article) (*article.Article, error) {
	article, err := s.requirements.Repo.Update(ctx, id, changes)
	if err != nil {
		return nil, ErrUpdating
	}

	return article, nil
}
