package dto

type CreateArticleDto struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Content     string `validate:"required"`
}
