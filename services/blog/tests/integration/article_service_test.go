package auth_integration_test

import (
	"context"
	"testing"

	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/domain/article"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ArticleServiceSuite struct {
	suite.Suite

	article *article.Article
}

func (s *ArticleServiceSuite) SetupSuite() {
	s.article = &article.Article{
		Title:       "Article 1",
		Description: "Article 1 description",
		Content:     "Content of the article 1",
		CoverImage:  "https://example.com/image.png",
	}
}

func (s *ArticleServiceSuite) TestA_Create() {
	ctx := context.TODO()

	ar, err := articleService.Create(ctx, s.article.Title, s.article.Description, s.article.Content, s.article.CoverImage)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), ar)
	require.Equal(s.T(), ar.Title, s.article.Title)
	require.Equal(s.T(), ar.Description, s.article.Description)
	require.Equal(s.T(), ar.Content, s.article.Content)
	require.Equal(s.T(), ar.CoverImage, s.article.CoverImage)

	s.article = ar
}
func (s *ArticleServiceSuite) TestB_Update() {
	ctx := context.TODO()

	changes := &article.Article{Title: "Article 1 Title Must Be Changed"}
	ar, err := articleService.Update(ctx, s.article.ID, changes)

	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), ar)
	require.Equal(s.T(), ar.Title, changes.Title)
	require.NotEqual(s.T(), ar.Title, s.article.Title)

	s.article = ar
}

func (s *ArticleServiceSuite) TestC_Find() {
	ctx := context.TODO()

	ar, err := articleService.Find(ctx, s.article.ID)

	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), ar)
	require.Equal(s.T(), ar.Title, s.article.Title)
	require.Equal(s.T(), ar.Description, s.article.Description)
	require.Equal(s.T(), ar.Content, s.article.Content)
	require.Equal(s.T(), ar.CoverImage, s.article.CoverImage)
}

func (s *ArticleServiceSuite) TestD_Search() {
	ctx := context.TODO()

	searchInput := "Article 1 Title Must Be Changed"
	articles, err := articleService.Search(ctx, searchInput)

	require.NoError(s.T(), err)
	require.Len(s.T(), articles, 1)
	require.Equal(s.T(), articles[0].Title, s.article.Title)
	require.Equal(s.T(), articles[0].Description, s.article.Description)
	require.Equal(s.T(), articles[0].Content, s.article.Content)
	require.Equal(s.T(), articles[0].CoverImage, s.article.CoverImage)
}

func TestArticleServiceSuite(t *testing.T) {
	t.Helper()
	suite.Run(t, new(ArticleServiceSuite))
}
