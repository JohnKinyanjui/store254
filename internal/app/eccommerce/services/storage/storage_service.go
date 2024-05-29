package storage_service

import "eccomerce_apis/internal/config"

func UploadFiles(cfg config.Config, data StorageData) ([]string, error) {
	return uploadFilesImages(cfg.Cloudinary, data.Images, data.ImageHeaders)
}
