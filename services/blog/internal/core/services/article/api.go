package article_service

import (
	"context"

	"github.com/tahadostifam/go-hexagonal-architecture/internal/core/domain/article"
)

type Api interface {
	Create(ctx context.Context, title, description, content, coverImage string) (*article.Article, error)
	Update(ctx context.Context, id int64, changes *article.Article) (*article.Article, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*article.Article, error)
	Search(ctx context.Context, title string) ([]article.Article, error)
}
