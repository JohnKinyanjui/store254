package main

import (
	public_router "eccomerce_apis/internal/app/public/rest"

	"eccomerce_apis/internal/config"

	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/cors"
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

	app := iris.New()
	app.Use(logger.New())
	app.Validator = validator.New()

	// CORS Configuration to accept from any HTTPS-verified domains and localhost.
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*"}, // Accept from any HTTPS verified domains
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true // you might want to implement more sophisticated checks here
		},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
	})
	app.UseRouter(crs)
	app.AllowMethods(iris.MethodOptions)

	cfg := config.Config{}
	cfg.Init()

	// Cookie configuration for SameSite and Secure attributes.
	cfg.App = app

	public_router.ConfigPublicAPIS(&cfg)

	port := "8001"
	host := fmt.Sprintf("0.0.0.0:%s", port)

	err := app.Listen(host, iris.WithPostMaxMemory(maxSize))
	if err != nil {
		panic(fmt.Sprint("unable to run http server: ", err.Error()))
	}
}
