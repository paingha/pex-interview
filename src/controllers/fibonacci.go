package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paingha/pex/src/lib"
)

func (env *Env) Next(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodPost).WithRequestPath(c.FullPath()).Info("API", "Next route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Next())
}

func (env *Env) Previous(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodPost).WithRequestPath(c.FullPath()).Info("API", "Previous route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Previous())
}

func (env *Env) Current(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodGet).WithRequestPath(c.FullPath()).Info("API", "Current route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Current())
}
