package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type TestConfig struct {
	Env        string `env:"ENV"`
	EnvDefault string `env:"ENV_DEFAULT" envDefault:"test"`
}

func cleanEnv() {
	os.Unsetenv("ENV")
	os.Unsetenv("ENV_DEFAULT")
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestLoadConfig(t *testing.T) {
	cleanEnv()

	t.Run("Should load config from environment variables", func(t *testing.T) {
		os.Setenv("ENV", "test")
		os.Setenv("ENV_DEFAULT", "test2")

		config := &TestConfig{}
		err := LoadConfig(config)

		assert.Nil(t, err)
		assert.Equal(t, "test", config.Env)
		assert.Equal(t, "test2", config.EnvDefault)
	})

	cleanEnv()

	t.Run("Should load config from environment variables with default values", func(t *testing.T) {
		os.Setenv("ENV", "test")

		config := &TestConfig{}
		err := LoadConfig(config)

		assert.Nil(t, err)
		assert.Equal(t, "test", config.Env)
		assert.Equal(t, "test", config.EnvDefault)
	})
}
