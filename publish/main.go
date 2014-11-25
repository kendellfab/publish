package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/interfaces"
	"github.com/kendellfab/publish/usecases"
	"log"
)

func main() {
	var configPath = flag.String("c", "config.toml", "Set the config path.")
	flag.Parse()

	var config domain.Config
	_, tomlErr := toml.DecodeFile(*configPath, &config)
	if tomlErr != nil {
		log.Fatal(tomlErr)
	}
	log.Println(config)

	db, dbErr := sql.Open("sqlite3", config.Sqlite)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	repoManager := usecases.RepoManager{}
	repoManager.CommentRepo = interfaces.NewDbCommentRepo(db)
	repoManager.UserRepo = interfaces.NewDbUserRepo(db)
	repoManager.CategoryRepo = interfaces.NewDbCategoryRepo(db)
	repoManager.PostRepo = interfaces.NewDbPostRepo(db, repoManager.UserRepo, repoManager.CategoryRepo)
	repoManager.ContactRepo = interfaces.NewDbContactRepo(db)
	repoManager.PageRepo = interfaces.NewDbPageRepo(db)

}
