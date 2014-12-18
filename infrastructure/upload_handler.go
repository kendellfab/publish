package infrastructure

import (
	"github.com/kendellfab/publish/usecases"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type UploadHandler struct {
	baseDir string
}

func NewUploadHandler(base string) usecases.FileRepo {
	uh := UploadHandler{baseDir: base}
	if err := os.MkdirAll(base, 0700); err != nil {
		log.Fatal(err)
	}
	return uh
}

func (uh UploadHandler) SaveFile(input *os.File) error {
	outfile, outErr := os.Create(filepath.Join(uh.baseDir, input.Name()))
	if outErr != nil {
		return outErr
	}
	defer outfile.Close()

	size, err := io.Copy(outfile, input)
	if err != nil {
		return err
	}
	log.Println("Uploaded: ", input.Name(), "Size:", size)
	return nil
}

func (uh UploadHandler) SaveMultipartFile(input *multipart.FileHeader) error {
	log.Println(input.Filename)
	infile, inErr := input.Open()
	if inErr != nil {
		return inErr
	}
	defer infile.Close()

	outfile, outErr := os.Create(filepath.Join(uh.baseDir, input.Filename))
	if outErr != nil {
		return outErr
	}
	defer outfile.Close()

	size, err := io.Copy(outfile, infile)
	if err != nil {
		return err
	}
	log.Println("Uploaded: ", input.Filename, "Size:", size)

	return nil
}

func (uh UploadHandler) DeleteFile(file string) error {
	path := filepath.Join(uh.baseDir, file)
	return os.Remove(path)
}

func (uh UploadHandler) ListFiles() ([]string, error) {
	return nil, nil
}
