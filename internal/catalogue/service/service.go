package service

import (
	"eventori/internal/catalogue"
	"eventori/pkg/helper"
	"log"
	"mime/multipart"
	"strings"
)

type uploader func(image *multipart.FileHeader) (string, error)

type catalogueService struct {
	repo   catalogue.Repository
	upload uploader
}

func New(r catalogue.Repository, u uploader) catalogue.Service {
	return &catalogueService{repo: r, upload: u}
}

func (cs *catalogueService) GetList() ([]catalogue.Core, error) {
	catalogueCores, err := cs.repo.GetList()
	if err != nil {
		log.Println(err)
		return nil, helper.SetError(500, "internal server error")
	}

	return catalogueCores, nil
}

func (cs *catalogueService) Create(catalogueCore catalogue.Core, image *multipart.FileHeader) error {
	if len(catalogueCore.ModelName) < 3 {
		log.Println("Model name must greater than 3")
		return helper.SetError(400, "bad request")
	}
	// image limit 5 mb
	if image.Size > (5 * 1024 * 1024) {
		log.Println("image out of sized")
		return helper.SetError(400, "bad request")
	}
	fileimg := strings.Split(image.Filename, ".")
	format := fileimg[len(fileimg)-1]
	if format != "png" && format != "jpg" && format != "jpeg" {
		log.Println("bad request because of format not png, jpg, or jpeg")
		return helper.SetError(400, "bad request")
	}

	imageURL, err := cs.upload(image)
	if err != nil {
		log.Println(err)
		return helper.SetError(500, "internal server error")
	}

	catalogueCore.ImageURL = imageURL
	if err := cs.repo.Create(catalogueCore); err != nil {
		log.Println(err)
		return helper.SetError(500, "internal server error")
	}

	return nil
}

func (cs *catalogueService) GetByID(modelID int) (catalogue.Core, error) {
	catalogueCore, err := cs.repo.GetByID(modelID)
	if err != nil {
		log.Println(err)
		return catalogue.Core{}, helper.SetError(500, "internal server error")
	}

	return catalogueCore, nil
}

func (cs *catalogueService) Update(modelID int, catalogueCore catalogue.Core, image *multipart.FileHeader) error {
	if len(catalogueCore.ModelName) > 0 && len(catalogueCore.ModelName) < 3 {
		log.Println("Model name must greater than 3")
		return helper.SetError(400, "bad request")
	}
	if image != nil {
		// image limit 5 mb
		if image.Size > (5 * 1024 * 1024) {
			log.Println("image out of sized")
			return helper.SetError(400, "bad request")
		}
		fileimg := strings.Split(image.Filename, ".")
		format := fileimg[len(fileimg)-1]
		if format != "png" && format != "jpg" && format != "jpeg" {
			log.Println("bad request because of format not png, jpg, or jpeg")
			return helper.SetError(400, "bad request")
		}

		imageURL, err := cs.upload(image)
		if err != nil {
			log.Println(err)
			return helper.SetError(500, "internal server error")
		}

		catalogueCore.ImageURL = imageURL
	}

	if err := cs.repo.Update(modelID, catalogueCore); err != nil {
		log.Println(err)
		return helper.SetError(500, "internal server error")
	}

	return nil
}

func (cs *catalogueService) Delete(modelID int) error {
	if err := cs.repo.Delete(modelID); err != nil {
		log.Println(err)
		return helper.SetError(500, "internal server error")
	}

	return nil
}
