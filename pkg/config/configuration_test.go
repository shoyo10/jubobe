package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWithSettingEnv(t *testing.T) {
	pwd := os.Getenv("PWD")
	t.Log(pwd)
	t.Setenv("CONFIG_DIR", pwd)
	t.Setenv("CONFIG_NAME", "testconfig")
	cfg, err := New()
	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	assert.True(t, cfg.Log.Debug)
	assert.Equal(t, "test", cfg.Log.AppID)
}

func TestNewWithoutSettingEnv(t *testing.T) {
	_, err := New()
	assert.ErrorContains(t, err, `Config File "app" Not Found in`)
}
