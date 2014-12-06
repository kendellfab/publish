package interfaces

import (
	"database/sql"
	"fmt"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
)

type DbUserRepo struct {
	db *sql.DB
}

func NewDbUserRepo(db *sql.DB) domain.UserRepo {
	userRepo := &DbUserRepo{db: db}
	userRepo.init()
	return userRepo
}

func (repo *DbUserRepo) init() {
	if _, err := repo.db.Exec(CREATE_USER); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbUserRepo) Store(user *domain.User) error {
	// repo.logger.Log("Storing user: " + user.Email + " " + user.Name)
	insertStmt := "INSERT INTO user(name, email, password, role) VALUES(?, ?, ?, ?)"
	res, err := repo.db.Exec(insertStmt, user.Name, user.Email, user.Password, user.Role)

	if err == nil {
		if id, idErr := res.LastInsertId(); idErr == nil {
			user.Id = id
		}
	}

	return err
}

func (repo *DbUserRepo) FindById(id string) (*domain.User, error) {
	var user domain.User
	row := repo.db.QueryRow("SELECT * FROM user WHERE id=?", id)
	scanErr := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
	if scanErr != nil {
		return nil, scanErr
	}

	return &user, nil
}

func (repo *DbUserRepo) FindByIdInt(id int64) (*domain.User, error) {
	var user domain.User
	row := repo.db.QueryRow("SELECT * FROM user WHERE id=?", id)
	scanErr := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
	if scanErr != nil {
		return nil, scanErr
	}

	return &user, nil
}

func (repo *DbUserRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	row := repo.db.QueryRow("SELECT * FROM user WHERE email=?", email)
	scanErr := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
	if scanErr != nil {
		return nil, scanErr
	}
	return &user, nil
}

func (repo *DbUserRepo) FindAdmin() (*[]domain.User, error) {
	rows, rowError := repo.db.Query("SELECT id, name, email, salt, role FROM user WHERE role=?", domain.Admin)
	if rowError != nil {
		return nil, rowError
	}
	users := scanUsers(rows)
	return &users, nil
}

func scanUsers(rows *sql.Rows) []domain.User {
	fmt.Println("Scanning...")
	users := make([]domain.User, 0)
	for {
		var user domain.User
		scanErr := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Role)
		if scanErr == nil {
			users = append(users, user)
		}
		if !rows.Next() {
			break
		}
	}
	return users
}
