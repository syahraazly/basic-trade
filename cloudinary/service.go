package cloudinary

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService interface {
	UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type service struct {
	cld *cloudinary.Cloudinary
}

func NewService() (*service, error) {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	if err != nil {
		return nil, err
	}

	return &service{cld: cld}, nil
}

func (s *service) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	useFilename := true
	uniqueFilename := true

	resp, err := s.cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       fileHeader.Filename,
		Folder:         "basic-trade",
		ResourceType:   "image",
		UseFilename:    &useFilename,
		UniqueFilename: &uniqueFilename,
	})

	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
