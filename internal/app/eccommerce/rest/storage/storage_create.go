package storage

import (
	storage_service "eccomerce_apis/internal/app/eccommerce/services/storage"

	"github.com/kataras/iris/v12"
)

func (api *StorageApi) uploadProductFiles(ctx iris.Context) {
	images, imageHeaders, err := ctx.FormFiles("images")
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error":      "unable to get images",
			"error_info": err.Error(),
		})
		return
	}

	attachments, attachmentHeaders, err := ctx.FormFiles("attachments")
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error":      "unable to get attachments",
			"error_info": err.Error(),
		})
		return
	}

	var data storage_service.StorageData
	data.Images = images
	data.ImageHeaders = imageHeaders
	data.Attachments = attachments
	data.AttachmentHeaders = attachmentHeaders

	links, err := storage_service.UploadFiles(*api.Config, data)
	if err != nil {
		_ = ctx.StopWithJSON(iris.StatusBadRequest, map[string]string{
			"error":      "unable to get attachments",
			"error_info": err.Error(),
		})
		return
	}

	_ = ctx.StopWithJSON(iris.StatusOK, map[string]interface{}{
		"images": links,
	})
}
