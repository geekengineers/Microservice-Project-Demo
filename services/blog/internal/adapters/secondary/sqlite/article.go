package sqlite_adapter

import (
	"context"
	"os/user"

	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/domain/article"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/ports"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(dialector gorm.Dialector) (ports.ArticleRepositorySecondaryPort, error) {
	db, err := GORM(dialector)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&user.User{})

	return &articleRepository{db}, nil
}

func (a *articleRepository) Create(ctx context.Context, articleModel *article.Article) (*article.Article, error) {
	tx := a.db.Model(&article.Article{}).Create(&articleModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return articleModel, nil
}

func (a *articleRepository) Find(ctx context.Context, id int64) (*article.Article, error) {
	var foundArticle *article.Article
	tx := a.db.Model(&article.Article{}).Where("id = ?", id).First(&foundArticle)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return foundArticle, nil
}

func (a *articleRepository) Search(ctx context.Context, title string) ([]article.Article, error) {
	var foundArticles []article.Article
	tx := a.db.Model(&article.Article{}).Where("title LIKE ?", title).Find(&foundArticles)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return foundArticles, nil
}

func (a *articleRepository) Update(ctx context.Context, id int64, changes *article.Article) (*article.Article, error) {
	ar, err := a.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	if changes.Title != "" {
		ar.Title = changes.Title
	}

	if changes.Content != "" {
		ar.Content = changes.Content
	}

	if changes.CoverImage != "" {
		ar.CoverImage = changes.CoverImage
	}

	if changes.Description != "" {
		ar.Description = changes.Description
	}

	if !changes.PublishedAt.IsZero() {
		ar.PublishedAt = changes.PublishedAt
	}

	tx := a.db.Save(ar)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ar, nil
}

func (a *articleRepository) Delete(ctx context.Context, id int64) error {
	var foundArticle *article.Article
	tx := a.db.Model(&article.Article{}).Where("id = ?", id).First(&foundArticle)
	if tx != nil {
		return tx.Error
	}

	tx = a.db.Unscoped().Delete(&foundArticle)
	if tx != nil {
		return tx.Error
	}

	return nil
}
