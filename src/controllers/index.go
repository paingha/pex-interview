package controllers

import (
	"github.com/paingha/pex/src/lib"
	"github.com/paingha/pex/src/models"
)

type Env struct {
	Db     *models.FibonacciDataStore
	Logger *lib.Log
}
