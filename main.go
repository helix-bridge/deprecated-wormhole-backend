package main

import (
	"fmt"
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/middlewares"
	"github.com/darwinia-network/link/observer"
	serverHttp "github.com/darwinia-network/link/server/routes/http"
	"github.com/darwinia-network/link/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/urfave/cli.v2"
	"net/http"
	"os"
	"runtime"
)

func main() {
	defer util.CloseDB()
	if err := setupApp().Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func setupApp() *cli.App {
	return &cli.App{
		Name:    "Darwinia-Dapp",
		Version: "0.1",
		Before: func(context *cli.Context) error {
			runtime.GOMAXPROCS(runtime.NumCPU())
			config.LoadConf()
			return nil
		},
		Action: func(c *cli.Context) error {
			util.GraceShutdown(&http.Server{Addr: ":5333", Handler: setupRouter()})
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "observer",
				Action: func(c *cli.Context) error {
					observer.Run()
					util.RunForever()
					return nil
				},
			},
		},
	}
}

func setupRouter() (server *gin.Engine) {
	server = gin.Default()
	server.Use(middlewares.CORS())
	serverHttp.Run(server.Group("/api"))
	return
}
