package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"api/api/middleware"
	"api/config"

	backendRoute "api/api/modules/backend/route"
)

func StartApi() *gin.Engine {
	gin.SetMode(config.Mode())
	if config.Mode() == config.ReleaseMode {
		gin.DisableConsoleColor()
	}

	r := initAPIEngine()

	backendRoute.Route(r)
	return r
}

func initAPIEngine() *gin.Engine {
	r := gin.New()

	r.HandleMethodNotAllowed = true
	r.RedirectTrailingSlash = false

	if config.IsDebugMode() {
		r.Use(gin.Logger())
	}

	r.Use(middleware.CORS())
	r.Use(middleware.SecurityHeaders())
	r.Use(middleware.RequestRate())

	return r
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL)
	sig := <-ch
	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("------exited--------", time.Since(now))
}
