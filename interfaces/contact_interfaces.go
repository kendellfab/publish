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
	if _, err := repo.db.Exec(CREATE_CONTACT); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
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
