package usecases

import (
	"mime/multipart"
	"os"
)

type FileRepo interface {
	SaveFile(input *os.File) error
	SaveMultipartFile(input *multipart.FileHeader) error
	DeleteFile(file string) error
	ListFiles() ([]os.FileInfo, error)
}
