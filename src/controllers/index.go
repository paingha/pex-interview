package controllers

import (
	"github.com/paingha/pex/src/lib"
	"github.com/paingha/pex/src/models"
)

// Env struct containing dependency injected fields.
type Env struct {
	Db     *models.FibonacciDataStore
	Logger *lib.Log
}
