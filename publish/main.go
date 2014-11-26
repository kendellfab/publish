package main

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/sessions"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/handlers/admin"
	"github.com/kendellfab/publish/handlers/front"
	"github.com/kendellfab/publish/interfaces"
	"github.com/kendellfab/publish/usecases"
	"log"
	"path/filepath"
)

func main() {
	var configPath = flag.String("c", "config.toml", "Set the config path.")
	flag.Parse()

	var config domain.Config
	_, tomlErr := toml.DecodeFile(*configPath, &config)
	if tomlErr != nil {
		log.Fatal(tomlErr)
	}

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

	adminRender := milo.NewDefaultRenderer(filepath.Join(config.AdminDir, "tpls"), false, nil)
	frontendRender := milo.NewDefaultRenderer(filepath.Join(config.ThemeDir, "tpls"), false, nil)
	store := sessions.NewCookieStore([]byte(config.SessionKeys[0]))

	adminBase := admin.NewAdminBase(adminRender, repoManager, store)
	adminGen := admin.NewAdminGeneral(&adminBase, repoManager)

	frontBase := front.NewFrontBase(frontendRender, repoManager)

	app := milo.NewMiloApp(milo.SetPort(config.Port))

	adminGen.RegisterRoutes(app)
	frontBase.RegisterRoutes(app)

	app.RouteAssetStripPrefix("/admin", config.AdminDir)
	app.RouteAsset("/css", config.ThemeDir)
	app.RouteAsset("/js", config.ThemeDir)

	app.Run()
}
