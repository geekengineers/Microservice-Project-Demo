package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigsDirPath(t *testing.T) {
	t.Log(ConfigsDirPath())
	t.Log(ProjectRootPath)
}

func TestDevelopmentConfig(t *testing.T) {
	os.Setenv("KAVKA_ENV", "development")

	config := Read()
	require.NotEmpty(t, config)
	require.NotZero(t, config.Grpc.Port)
	require.Equal(t, CurrentEnv, Development)

	t.Log(config)
}
