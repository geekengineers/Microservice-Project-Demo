package auth_integration_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ArticleServiceSuite struct {
	suite.Suite
}

func (s *ArticleServiceSuite) SetupSuite() {

}

func (s *ArticleServiceSuite) TestA_Login() {
}

func TestArticleServiceSuite(t *testing.T) {
	t.Helper()
	suite.Run(t, new(ArticleServiceSuite))
}
