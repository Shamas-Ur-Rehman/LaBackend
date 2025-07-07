package utils

import (
	"Laorgaincs/config"
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Reset the file reader (in case it's been read before)
	file.Seek(0, 0)

	uploadResult, err := config.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fmt.Sprintf("product_images/%d_%s", time.Now().Unix(), fileHeader.Filename),
		Folder:   "product_images",
	})

	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
