package interfaces

import (
	"database/sql"
	"github.com/kendellfab/publish/domain"
	"log"
	"strings"
	"time"
)

type DbCommentRepo struct {
	db *sql.DB
}

func NewDbCommentRepo(db *sql.DB) domain.CommentRepo {
	commentRepo := &DbCommentRepo{db: db}
	commentRepo.init()
	return commentRepo
}

func (repo *DbCommentRepo) init() {
	exec := `CREATE TABLE comment (
"id" INTEGER NOT NULL,
"page" VARCHAR(256) NOT NULL,
"username" VARCHAR(64),
"email" VARCHAR(64),
"date" VARCHAR(64),
"content" TEXT,
"approved" INTEGER DEFAULT (0),
"day" INTEGER, 
"month" INTEGER, 
"year" INTEGER)`

	if _, err := repo.db.Exec(exec); err != nil && !strings.Contains(err.Error(), domain.ALREADY_EXISTS) {
		log.Fatal(err)
	}
}

func (repo *DbCommentRepo) Store(comment *domain.Comment) error {
	day, month, year := domain.DateComponents(comment.Date)
	createdStr := domain.SerializeDate(comment.Date)
	_, err := repo.db.Exec("INSERT INTO comment(page, username, email, date, content, approved, day, month, year) VALUES(?, ?, ?, ?, ?, 0, ?, ?, ?)", comment.Page, comment.Username, comment.Email, createdStr, comment.Content, day, month, year)
	return err
}

func (repo *DbCommentRepo) FindById(id int) (*domain.Comment, error) {
	return nil, nil
}

func (repo *DbCommentRepo) FindByPage(page string) (*[]domain.Comment, error) {
	sql := "SELECT * FROM comment WHERE page=?"
	rows, qError := repo.db.Query(sql, page)
	if qError != nil {
		return nil, qError
	}
	comments := scanComments(rows)
	return &comments, nil
}

func (repo *DbCommentRepo) FindUnapprovedComments() (*[]domain.Comment, error) {
	sql := "SELECT * FROM comment WHERE approved = 0"
	rows, qError := repo.db.Query(sql)
	if qError != nil {
		return nil, qError
	}
	comments := scanComments(rows)
	return &comments, nil
}

func (repo *DbCommentRepo) ApproveComment(id int) error {
	sql := "UPDATE comment SET approved = 1 WHERE id = ?"
	_, err := repo.db.Exec(sql, id)
	return err
}

func scanComments(rows *sql.Rows) []domain.Comment {
	comments := make([]domain.Comment, 0)

	for {
		var comment domain.Comment
		var approved int
		var date string
		scanErr := rows.Scan(&comment.Id, &comment.Page, &comment.Username, &comment.Email, &date, &comment.Content, &approved)
		if approved == 0 {
			comment.Approved = false
		} else {
			comment.Approved = true
		}
		if scanErr == nil {
			comment.Date, _ = time.Parse(time.RFC3339, date)
			comments = append(comments, comment)
		}

		if !rows.Next() {
			break
		}
	}
	return comments
}
