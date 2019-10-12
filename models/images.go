package models

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Image is not stored in the database
type Image struct {
	GalleryID uint
	Filename  string
}

func (i *Image) Path() string {
	return "/" + i.RelativePath()
}

func (i *Image) RelativePath() string {
	return fmt.Sprintf("images/galleries/%v/%v", i.GalleryID, i.Filename)
}

type ImageService interface {
	Create(galleryID uint, r io.ReadCloser, filename string) error
	Delete(image *Image) error
	ByGalleryID(galleryID uint) ([]Image, error)
}

func NewImageService() ImageService {
	return &imageService{}
}

type imageService struct{}

func (is *imageService) Create(galleryID uint, r io.ReadCloser, filename string) error {
	defer r.Close()
	path, err := is.mkImagePath(galleryID)
	if err != nil {
		return err
	}

	dst, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, r)
	if err != nil {
		return err
	}
	return nil
}

func (is *imageService) Delete(i *Image) error {
	path := i.RelativePath()
	if strings.Contains(path, "/") {
		path = strings.ReplaceAll(path, "/", "\\")
	}
	return os.Remove(path)
}

func (is *imageService) ByGalleryID(galleryID uint) ([]Image, error) {
	path := is.imagePath(galleryID)
	files, err := filepath.Glob(path + "*")
	if err != nil {
		return nil, err
	}
	ret := make([]Image, len(files))
	for i := range files {
		f := files[i]
		if strings.Contains(f, "\\") {
			files[i] = strings.ReplaceAll(f, "\\", "/")
		}
		files[i] = strings.Replace(files[i], path, "", 1)
		ret[i] = Image{
			Filename:  files[i],
			GalleryID: galleryID,
		}
	}
	return ret, nil
}

func (is *imageService) imagePath(galleryID uint) string {
	return fmt.Sprintf("images/galleries/%v/", galleryID)
}
func (is *imageService) mkImagePath(galleryID uint) (string, error) {
	galleryPath := is.imagePath(galleryID)
	err := os.MkdirAll(galleryPath, 0755)
	if err != nil {
		return "", err
	}
	return galleryPath, nil
}
