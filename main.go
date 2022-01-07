package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paingha/pex/src/config"
	"github.com/paingha/pex/src/lib"
	"github.com/paingha/pex/src/models"
	"github.com/paingha/pex/src/routes"
)

func main() {
	logger := lib.NewLogger()
	systemConfig, err := config.InitSystemConfig("")
	if err != nil {
		logger.Fatal("API", lib.FormatErrorMessage("An Error occured while reading config file", err), lib.FileZone())
	}
	db := models.NewFibonacciDataStore()
	env := config.NewENV(db, logger)
	gin.SetMode(systemConfig.Environment)
	r := routes.SetupRouter(env)
	srv := &http.Server{
		Addr:    systemConfig.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("API", lib.FormatErrorMessage("An Error occured while starting the server", err), lib.FileZone())
		}
	}()

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)
	<-close

	logger.Info("API", "Shutting Down Server", lib.FileZone())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("API", lib.FormatErrorMessage("Server Shutdown:", err), lib.FileZone())
	}
	select {
	case <-ctx.Done():
		logger.Info("API", "timeout of 5 seconds.", lib.FileZone())
	}
	logger.Info("API", "Server exiting", lib.FileZone())
}
