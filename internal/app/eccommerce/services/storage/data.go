package storage_service

import "mime/multipart"

type StorageData struct {
	Images            []multipart.File        `json:"-"`
	ImageHeaders      []*multipart.FileHeader `json:"-"`
	Attachments       []multipart.File        `json:"-"`
	AttachmentHeaders []*multipart.FileHeader `json:"-"`
}
