package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/sessions"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/handlers/admin"
	"github.com/kendellfab/publish/handlers/front"
	"github.com/kendellfab/publish/infrastructure"
	"github.com/kendellfab/publish/interfaces"
	"github.com/kendellfab/publish/usecases"
	"log"
	"path/filepath"
)

func main() {
	log.Println("Starting publish... DB:", infrastructure.CurrentDb)
	var configPath = flag.String("c", "config.toml", "Set the config path.")
	flag.Parse()

	var config domain.Config
	_, tomlErr := toml.DecodeFile(*configPath, &config)
	if tomlErr != nil {
		log.Fatal(tomlErr)
	}

	db, dbErr := infrastructure.ConnectDb(&config)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer db.Close()
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	// viewDb, vdbErr := infrastructure.ConnectDb(&config)
	// if vdbErr != nil {
	// 	log.Fatal(vdbErr)
	// }
	// defer viewDb.Close()

	repoManager := usecases.RepoManager{}
	repoManager.CommentRepo = interfaces.NewDbCommentRepo(db)
	repoManager.UserRepo = interfaces.NewDbUserRepo(db)
	repoManager.CategoryRepo = interfaces.NewDbCategoryRepo(db)
	repoManager.PostRepo = interfaces.NewDbPostRepo(db, repoManager.UserRepo, repoManager.CategoryRepo)
	repoManager.ContactRepo = interfaces.NewDbContactRepo(db)
	repoManager.PageRepo = interfaces.NewDbPageRepo(db)
	repoManager.PayloadRepo = interfaces.NewPayloadRepo(config, repoManager.CategoryRepo, repoManager.PostRepo)
	repoManager.ViewRepo = interfaces.NewDbViewRepo(db)

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

	frontBase := front.NewFrontBase(frontendRender, repoManager, config)
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
