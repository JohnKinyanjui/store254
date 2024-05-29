package config

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"k8s.io/client-go/kubernetes"

	"github.com/go-pg/pg/v10"
	"github.com/kataras/iris/v12"
)

type Config struct {
	App        *iris.Application
	DB         *pg.DB
	BDB        *bun.DB
	RDB        *redis.Client
	Cloudinary *cloudinary.Cloudinary
	Fcm        *firebase.App
	Messaging  *messaging.Client
	K8         *kubernetes.Clientset
}

func (cfg *Config) Init() {
	cfg.getFirebaseClient()
	cfg.getDatabases()
	cfg.getStorages()
	cfg.getRedisClient()
	cfg.getKubernetesClient()
}
