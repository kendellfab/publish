package usecases

import (
	"github.com/kendellfab/publish/domain"
	"mime/multipart"
	"os"
)

type FileRepo interface {
	SaveFile(input *os.File) error
	SaveMultipartFile(input *multipart.FileHeader) error
	DeleteFile(file string) error
	ListFiles() ([]*domain.UploadNode, error)
}
