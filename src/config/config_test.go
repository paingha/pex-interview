package config_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/paingha/pex/src/config"
	"github.com/paingha/pex/src/lib"
	"github.com/paingha/pex/src/models"
	"github.com/stretchr/testify/assert"
)

func TestNewSystemConfig(t *testing.T) {
	sysConfig := config.NewSystemConfig()
	assert.NotEmpty(t, sysConfig)
}

func TestNewENV(t *testing.T) {
	logger := lib.NewLogger()
	db := models.NewFibonacciDataStore()
	env := config.NewENV(db, logger)
	assert.NotEmpty(t, env)
}

func TestInitSystemConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "pex_*")
	assert.NoError(t, err)
	defer os.RemoveAll(dir)
	fileContents := []byte("env: debug\nport: \":8080\"")
	filePath := fmt.Sprintf("/%s/%s", dir, "config.yaml")
	err = os.WriteFile(filePath, fileContents, 0644)
	assert.NoError(t, err)
	t.Log(filePath)
	sysConfig, err := config.InitSystemConfig(filePath)
	assert.NoError(t, err)
	assert.NotEmpty(t, sysConfig)
}
