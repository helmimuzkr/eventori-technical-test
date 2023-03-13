package cloud

import (
	"context"
	"eventori/config"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type cloudinaryInstance struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinary(conf *config.Config) *cloudinaryInstance {
	cld, err := cloudinary.NewFromParams(conf.CloudinaryName, conf.CloudinaryApiKey, conf.CloudinaryApiSecret)
	if err != nil {
		log.Println("init cloudinary gagal", err)
		return nil
	}

	return &cloudinaryInstance{cld: cld}
}

func (ci *cloudinaryInstance) Upload(file *multipart.FileHeader) (string, error) {
	// Open file header
	src, _ := file.Open()
	defer src.Close()

	// Upload file using formated public id
	publicID := fmt.Sprintf("%d-%s", int(file.Size), time.Now().Format("20060102-150405")) // Format  "file_size-(YY-MM-DD)-(hh-mm-ss)""
	uploadResult, err := ci.cld.Upload.Upload(
		context.Background(),
		src,
		uploader.UploadParams{
			PublicID: publicID,
			Folder:   "file",
		})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
