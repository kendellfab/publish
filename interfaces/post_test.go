package interfaces

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kendellfab/publish/domain"
	"log"
	"testing"
)

var posts []*domain.Post

func BenchmarkLatestPosts(b *testing.B) {
	db, err := sql.Open("mysql", "publish:publish@/publish")
	if err != nil {
		log.Println(err)
		return
	}
	userRepo := NewDbUserRepo(db)
	catRepo := NewDbCategoryRepo(db)
	postRepo := NewDbPostRepo(db, userRepo, catRepo)

	var ps []*domain.Post
	b.ResetTimer()
	// log.Println("Iterations:", b.N)
	for i := 0; i < b.N; i++ {
		items, _ := postRepo.FindPublished(0, 10)
		// log.Println(len(items))
		ps = items
	}

	posts = ps
}
