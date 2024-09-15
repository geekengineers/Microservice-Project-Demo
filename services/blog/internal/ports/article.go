package ports

import (
	"context"

	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/domain/article"
)

type ArticleRepositorySecondaryPort interface {
	Create(ctx context.Context, article *article.Article) (*article.Article, error)
	Update(ctx context.Context, id int64, changes *article.Article) (*article.Article, error)
	Find(ctx context.Context, id int64) (*article.Article, error)
	Search(ctx context.Context, title string) ([]article.Article, error)
	Delete(ctx context.Context, id int64) error
}
