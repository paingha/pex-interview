package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paingha/pex/src/lib"
)

// Next controllers method returns the next fibonacci number.
// @Summary Next Fibonacci number
// @Description get the next fibonacci number in the sequence
// @Accept  json
// @Produce  json
// @Success 200 {integer} 2
// @Router /fibonacci/next [post]
func (env *Env) Next(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodPost).WithRequestPath(c.FullPath()).Info("API", "Next route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Next())
}

// Previous controllers method returns the previous fibonacci number.
// @Summary Previous Fibonacci number
// @Description get the previous fibonacci number in the sequence
// @Accept  json
// @Produce  json
// @Success 200 {integer} 1
// @Router /fibonacci/previous [post]
func (env *Env) Previous(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodPost).WithRequestPath(c.FullPath()).Info("API", "Previous route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Previous())
}

// Current controllers method returns the current fibonacci number.
// @Summary Current Fibonacci number
// @Description get the current fibonacci number in the sequence
// @Accept  json
// @Produce  json
// @Success 200 {integer} 0
// @Router /fibonacci/current [get]
func (env *Env) Current(c *gin.Context) {
	env.Logger.WithRequestID().WithRequestMethod(http.MethodGet).WithRequestPath(c.FullPath()).Info("API", "Current route", lib.FileZone())
	c.JSON(http.StatusOK, env.Db.Current())
}
