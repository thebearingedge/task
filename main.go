package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thebearingedge/task/application"
	"github.com/thebearingedge/task/gateway"
	"github.com/thebearingedge/task/log"
	"github.com/thebearingedge/task/server"
)

func main() {
	names := gateway.NewNameGateway(http.DefaultClient, os.Getenv("NAMES_SERVICE_BASE_URL"))
	jokes := gateway.NewJokeGateway(http.DefaultClient, os.Getenv("JOKES_SERVICE_BASE_URL"))
	app := application.NewApplication(names, jokes)
	logger := log.NewLogger()
	server := server.NewServer(app, logger)
	server.Use(
		gin.LoggerWithWriter(gin.DefaultWriter),
		gin.Recovery(),
	)
	gin.SetMode(gin.ReleaseMode)
	if err := server.Run(os.Getenv("LISTEN_ADDRESS")); err != nil {
		panic(err)
	}
}
