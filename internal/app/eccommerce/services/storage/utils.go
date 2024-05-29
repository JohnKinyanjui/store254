package storage_service

import (
	"context"
	"eccomerce_apis/internal/app/eccommerce/models"
	"eccomerce_apis/pkg/middlewares"
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func uploadFilesImages(cl *cloudinary.Cloudinary, files []multipart.File, fileHeaders []*multipart.FileHeader) ([]string, error) {
	var images []string

	if len(files) == 0 {
		return nil, errors.New("no image was added for the product")
	}

	for i, file := range files {
		var link string
		if os.Getenv("STORAGE") == "google_storage" {
			l, err := middlewares.GoogleStorage().Path("images").Upload(fileHeaders[i], &file)
			if err != nil {
				return nil, fmt.Errorf("unable to upload to filestac reason: %s", err.Error())
			}
			link = l
		} else {
			resp, err := cl.Upload.Upload(context.Background(), file, uploader.UploadParams{PublicID: uuid.New().String()})
			if err != nil {
				return nil, err
			}

			link = resp.URL
		}

		images = append(images, link)
	}

	if len(images) == 0 {
		return nil, errors.New("no image was added for the product")
	}

	return images, nil
}

func uploadAttachments(files []multipart.File, fileHeaders []*multipart.FileHeader) ([]models.ProductAttachment, error) {
	var attachments []models.ProductAttachment

	for i, ffile := range files {
		link, err := middlewares.GoogleStorage().Path("attachments").Upload(fileHeaders[i], &ffile)
		if err != nil {
			return nil, fmt.Errorf("unable to upload to filestac reason: %s", err.Error())
		}

		attachmentType := fileHeaders[i].Header.Get("Content-Type")

		attachments = append(attachments, models.ProductAttachment{
			Name: fileHeaders[i].Filename,
			Type: attachmentType,
			Link: link,
		})
	}

	return attachments, nil
}
