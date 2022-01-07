package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paingha/pex/src/lib"
)

// Next controllers method returns the next fibonacci number.
func (env *Env) Next(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodPost).WithRequestPath(c.FullPath()).Info("API", "Next route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Next())
}

// Previous controllers method returns the previous fibonacci number.
func (env *Env) Previous(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodPost).WithRequestPath(c.FullPath()).Info("API", "Previous route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Previous())
}

// Current controllers method returns the current fibonacci number.
func (env *Env) Current(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodGet).WithRequestPath(c.FullPath()).Info("API", "Current route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Current())
}
