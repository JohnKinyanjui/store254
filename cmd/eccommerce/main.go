package main

import (
	eccommerce_crons "eccomerce_apis/internal/app/eccommerce/crons"
	private_router "eccomerce_apis/internal/app/eccommerce/rest"
	"eccomerce_apis/internal/config"

	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

const maxSize = 500 * iris.MB

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading environment variables file")
			return
		}
	}

	app := iris.Default()
	app.Use(logger.New())
	app.RegisterView(iris.HTML("./templates/pages", ".html"))

	app.Validator = validator.New()

	cfg := config.Config{}
	cfg.App = app
	cfg.Init()

	private_router.ConfigAPIS(&cfg)
	// [ start cron jobs]
	eccommerce_crons.StartAllCrons(&cfg)

	// [start http server with websockets]

	// wst := neffos.New(websocket.DefaultGobwasUpgrader, ws.WsEvents(&cfg))

	// wst.OnConnect = wsConfig.OnConnect(cfg.DB)
	// wst.OnDisconnect = wsConfig.OnDisconnect

	// app.Get("/ws", websocket.Handler(wst))
	port := os.Getenv("PORT")

	host := fmt.Sprintf("0.0.0.0:%s", port)

	err := app.Listen(host, iris.WithPostMaxMemory(maxSize))
	if err != nil {
		panic(fmt.Sprint("unable to run http server: ", err.Error()))
	}

}
