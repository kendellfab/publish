package main

import (
	_ "github.com/go-sql-driver/mysql"
	// _ "code.google.com/p/go-sqlite/go1/sqlite3"
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

	// db, dbErr := sql.Open("sqlite3", config.Sqlite)
	db, dbErr := sql.Open("mysql", "publish:publish@/publish")
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
	adminRender.RegisterTemplateFunc("fmt_date", usecases.FormatDate)
	adminRender.RegisterTemplateFunc("fmt_bool", usecases.FormatBool)
	adminRender.RegisterTemplateFunc("rend_md", usecases.RenderMarkdown)

	frontendRender := milo.NewDefaultRenderer(filepath.Join(config.ThemeDir, "tpls"), false, nil)
	frontendRender.RegisterTemplateFunc("fmt_date", usecases.FormatDate)
	frontendRender.RegisterTemplateFunc("fmt_bool", usecases.FormatBool)
	frontendRender.RegisterTemplateFunc("rend_md", usecases.RenderMarkdown)

	store := sessions.NewCookieStore([]byte(config.SessionKeys[0]))

	adminBase := admin.NewAdminBase(adminRender, repoManager, store)
	adminGen := admin.NewAdminGeneral(&adminBase, repoManager)
	adminPosts := admin.NewAdminPost(&adminBase, repoManager)
	adminCat := admin.NewAdminCat(&adminBase, repoManager)

	frontBase := front.NewFrontBase(frontendRender, repoManager)
	frontPosts := front.NewFrontPosts(&frontBase)
	frontCategories := front.NewFrontCategories(&frontBase)

	app := milo.NewMiloApp(milo.SetPort(config.Port))

	adminGen.RegisterRoutes(app)
	adminPosts.RegisterRoutes(app)
	adminCat.RegisterRoutes(app)
	frontBase.RegisterRoutes(app)
	frontPosts.RegisterRoutes(app)
	frontCategories.RegisterRoutes(app)

	app.RouteAssetStripPrefix("/admin", config.AdminDir)
	app.RouteAsset("/css", config.ThemeDir)
	app.RouteAsset("/js", config.ThemeDir)

	app.Run()
}
