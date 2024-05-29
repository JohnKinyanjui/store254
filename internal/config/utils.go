package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-pg/pg/v10"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"google.golang.org/api/option"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func (c *Config) getGithubClient() {
	// itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 1, 99, "./internal/config/github.pem")

	// Or for endpoints that require JWT authentication

}

func (c *Config) getFirebaseClient() {
	opt := option.WithCredentialsFile("./internal/config/fcm.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}

	messaging, err := app.Messaging(context.Background())
	if err != nil {
		panic(fmt.Sprintf("error initializing messaging app: %v", err))
	}

	c.Fcm = app
	c.Messaging = messaging
}

func (c *Config) getDatabases() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	c.BDB = bun.NewDB(sqldb, pgdialect.New())

	opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	if db == nil {
		log.Fatalf("database not connected")
	}

	c.DB = db
}

func (c *Config) getStorages() {
	CLOUDINARY_CLOUD_NAME := os.Getenv("CLOUDINARY_CLOUD_NAME")
	CLOUDINARY_API_KEY := os.Getenv("CLOUDINARY_API_KEY")
	CLOUDINARY_API_SECRET := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(CLOUDINARY_CLOUD_NAME, CLOUDINARY_API_KEY, CLOUDINARY_API_SECRET)
	if err != nil {
		panic(err.Error())
	}

	c.Cloudinary = cld
}

func (c *Config) getKubernetesClient() {
	kubeconfig := "./internal/config/kubernetes.yaml"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		// panic(err.Error())
	}

	// c.K8 = clientset

}

func (c *Config) getRedisClient() {
	opt, err := redis.ParseURL("redis://default:Klp6LDake6jBlcfDnpldJ4D1DKN2K2i4@roundhouse.proxy.rlwy.net:27514")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	c.RDB = rdb
}
