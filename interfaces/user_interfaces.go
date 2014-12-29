package interfaces

import (
	"database/sql"
	"fmt"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
)

type DbUserRepo struct {
	db    *sql.DB
	cache *AuthorCache
}

func NewDbUserRepo(db *sql.DB) domain.UserRepo {
	userRepo := &DbUserRepo{db: db}
	ac, err := NewAuthorCache(25)
	if err != nil {
		log.Fatal(err)
	}
	userRepo.cache = ac
	userRepo.init()
	return userRepo
}

func (repo *DbUserRepo) init() {
	if _, err := repo.db.Exec(CREATE_USER); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbUserRepo) Store(user *domain.User) error {
	insertStmt := "INSERT INTO user(name, email, hash, password, bio, token, role) VALUES(?, ?, ?, ?)"
	res, err := repo.db.Exec(insertStmt, user.Name, user.Email, user.Hash, user.Password, user.Bio, user.Token, user.Role)

	if err == nil {
		if id, idErr := res.LastInsertId(); idErr == nil {
			user.Id = id
		}
	}

	return err
}

func (repo *DbUserRepo) Update(user *domain.User) error {
	up := "UPDATE user SET name = ?, email = ?, hash = ?, bio = ?, token = ? WHERE id = ?;"
	_, err := repo.db.Exec(up, user.Name, user.Email, user.Hash, user.Bio, user.Token, user.Id)
	return err
}

func (repo *DbUserRepo) FindById(id string) (*domain.User, error) {
	if user, ok := repo.cache.Get(id); ok {
		return user, nil
	}
	row := repo.db.QueryRow("SELECT id, name, email, hash, password, bio, token, role FROM user WHERE id=?", id)
	user, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}
	repo.cache.Add(id, user)
	return user, nil
}

func (repo *DbUserRepo) FindByIdInt(id int64) (*domain.User, error) {
	if user, ok := repo.cache.Get(fmt.Sprintf("%d", id)); ok {
		return user, nil
	}
	row := repo.db.QueryRow("SELECT id, name, email, hash, password, bio, token, role FROM user WHERE id=?", id)
	user, err := repo.scanRow(row)
	if err != nil {
		return nil, err
	}
	repo.cache.Add(fmt.Sprintf("%d", id), user)
	return user, nil
}

func (repo *DbUserRepo) FindByEmail(email string) (*domain.User, error) {
	row := repo.db.QueryRow("SELECT id, name, email, hash, password, bio, token, role FROM user WHERE email=?", email)
	return repo.scanRow(row)
}

func (repo *DbUserRepo) FindAdmin() (*[]domain.User, error) {
	rows, rowError := repo.db.Query("SELECT id, name, email, salt, role FROM user WHERE role=?", domain.Admin)
	if rowError != nil {
		return nil, rowError
	}
	users := repo.scanUsers(rows)
	return &users, nil
}

func (repo *DbUserRepo) UpdatePassword(userId, password string) error {
	up := "UPDATE user SET password = ? WHERE id = ?;"
	_, err := repo.db.Exec(up, password, userId)
	return err
}

func (repo *DbUserRepo) scanRow(row *sql.Row) (*domain.User, error) {
	var user domain.User
	scanErr := row.Scan(&user.Id, &user.Name, &user.Email, &user.Hash, &user.Password, &user.Bio, &user.Token, &user.Role)
	if scanErr != nil {
		return nil, scanErr
	}
	return &user, nil
}

func (repo *DbUserRepo) scanUsers(rows *sql.Rows) []domain.User {
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
