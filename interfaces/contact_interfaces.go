package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
)

type DbContactRepo struct {
	db *sql.DB
}

func NewDbContactRepo(db *sql.DB) domain.ContactRepo {
	contactRepo := &DbContactRepo{db: db}
	contactRepo.init()
	return contactRepo
}

func (repo *DbContactRepo) init() {
	exec := `CREATE TABLE contact (
"id" INTEGER NOT NULL,
"name" TEXT NOT NULL,
"email" TEXT NOT NULL,
"message" TEXT NOT NULL,
"read" INTEGER)`

	if _, err := repo.db.Exec(exec); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbContactRepo) Store(contact *domain.ContactForm) error {
	return nil
}

func (repo *DbContactRepo) GetAll() (*[]domain.ContactForm, error) {
	return nil, nil
}

func (repo *DbContactRepo) Delete(id int) error {
	return nil
}
