package storage

import (
	"eccomerce_apis/internal/config"

	"github.com/kataras/iris/v12"
)

type StorageApi struct {
	Config *config.Config
}

func (api *StorageApi) Configure(r iris.Party) {
	r.Post("/files", api.uploadProductFiles)
}
