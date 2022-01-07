package lib_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/paingha/pex/src/lib"
	"github.com/stretchr/testify/assert"
)

func TestWithRequestMethod(t *testing.T) {
	logger := lib.NewLogger()
	logger.WithRequestMethod(http.MethodPost)
	assert.Equal(t, http.MethodPost, logger.RequestMethod)
}

func TestWithRequestID(t *testing.T) {
	logger := lib.NewLogger()
	logger.WithRequestID()
	assert.NotEmpty(t, logger.RequestID)
}

func TestNewRequestID(t *testing.T) {
	requestID := lib.NewRequestID()
	assert.NotEmpty(t, requestID)
}

func TestWithRequestPath(t *testing.T) {
	logger := lib.NewLogger()
	logger.WithRequestPath("/api/test")
	assert.NotEmpty(t, logger.RequestPath)
	assert.Equal(t, logger.RequestPath, "/api/test")
}

func TestNewLogger(t *testing.T) {
	logger := lib.NewLogger()
	assert.NotEmpty(t, logger)
}

func TestFormatErrorMessage(t *testing.T) {
	testErr := fmt.Errorf("this is a test error")
	expectedErr := fmt.Sprintf("%s - %v", "Test err", testErr)
	errMsg := lib.FormatErrorMessage("Test err", testErr)
	assert.Equal(t, expectedErr, errMsg)
}
