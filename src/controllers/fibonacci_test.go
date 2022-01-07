package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paingha/pex/src/config"
	"github.com/paingha/pex/src/lib"
	"github.com/paingha/pex/src/models"
	"github.com/paingha/pex/src/routes"
	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	logger := lib.NewLogger()
	db := models.NewFibonacciDataStore()
	env := config.NewENV(db, logger)
	r := routes.SetupRouter(env)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, routes.FibonacciRootRoute+routes.NextFibonacciNumberPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1", w.Body.String())
}

func TestPrevious(t *testing.T) {
	logger := lib.NewLogger()
	db := models.NewFibonacciDataStore()
	env := config.NewENV(db, logger)
	r := routes.SetupRouter(env)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, routes.FibonacciRootRoute+routes.PreviousFibonacciNumberPath, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "0", w.Body.String())
}

func TestCurrent(t *testing.T) {
	logger := lib.NewLogger()
	db := models.NewFibonacciDataStore()
	env := config.NewENV(db, logger)
	r := routes.SetupRouter(env)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, routes.FibonacciRootRoute+routes.CurrentFibonacciNumberPath, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "0", w.Body.String())
}
