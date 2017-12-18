package utils

import (
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func HandleImage(file io.Reader, header *multipart.FileHeader, baseUrl string) (string, string, error) {
	fileName := GetFileName(header.Filename)
	img, err := GetImageFrom(file)
	if err != nil {
		return "", "", err
	}
	out, err := SaveImage(baseUrl, fileName)
	if err != nil {
		return "", "", err
	}
	defer out.Close()
	err = ToJpeg(out, img)
	filePath := (strings.Join([]string{baseUrl, fileName + ".jpg"}, "/"))
	return fileName, filePath, err
}

func GetFileName(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func GetImageFrom(file io.Reader) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return img, err
}

func SaveImage(baseUrl string, fileName string) (*os.File, error) {
	out, err := os.Create(strings.Join([]string{".", baseUrl, fileName + ".jpg"}, "/"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return out, err
}

func ToJpeg(file *os.File, img image.Image) (err error) {
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal(err)
	}
	return
}
