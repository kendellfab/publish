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
	var configPrefix = flag.String("p", domain.PUBLISH_PREFIX, "Set the prefix to be used in environment variables.")
	flag.Parse()

	var config domain.Config
	_, tomlErr := toml.DecodeFile(*configPath, &config)
	if tomlErr != nil {
		log.Fatal(tomlErr)
	}

	config.Email = domain.EmailConfigEnvironmentOverride(*configPrefix, config.Email)

	db := infrastructure.ConnectDb(&config)
	defer db.Close()

	repoManager := usecases.RepoManager{}
	repoManager.CommentRepo = interfaces.NewDbCommentRepo(db)
	repoManager.UserRepo = interfaces.NewDbUserRepo(db)
	repoManager.CategoryRepo = interfaces.NewDbCategoryRepo(db)
	repoManager.PostRepo = interfaces.NewDbPostRepo(db, repoManager.UserRepo, repoManager.CategoryRepo)
	repoManager.ContactRepo = interfaces.NewDbContactRepo(db)
	repoManager.PageRepo = interfaces.NewDbPageRepo(db)
	repoManager.PayloadRepo = interfaces.NewPayloadRepo(config, repoManager.CategoryRepo, repoManager.PostRepo)
	repoManager.ViewRepo = interfaces.NewDbViewRepo(db)
	repoManager.ResetRepo = interfaces.NewDbResetRepo(db)
	repoManager.SeriesRepo = interfaces.NewDbSeriesRepo(db, repoManager.PostRepo)
	repoManager.UploadRepo = infrastructure.NewUploadHandler(config.UploadDir)

	adminRender := milo.NewRenderer(filepath.Join(config.AdminDir, "tpls"), config.CacheTpls, nil)
	adminRender.RegisterTemplateFunc("fmt_date", usecases.FormatDate)
	adminRender.RegisterTemplateFunc("fmt_bool", usecases.FormatBool)
	adminRender.RegisterTemplateFunc("rend_md", usecases.RenderMarkdown)

	msgRend := milo.NewMsgRender(filepath.Join(config.AdminDir, "msgs"))
	emailMessenger := infrastructure.NewEmailMessenger(config.Email)

	frontendRender := milo.NewRenderer(filepath.Join(config.ThemeDir, "tpls"), config.CacheTpls, nil)
	frontendRender.RegisterTemplateFunc("fmt_date", usecases.FormatDate)
	frontendRender.RegisterTemplateFunc("fmt_bool", usecases.FormatBool)
	frontendRender.RegisterTemplateFunc("rend_md", usecases.RenderMarkdown)

	store := sessions.NewCookieStore([]byte(config.SessionKeys[0]))

	adminBase := admin.NewAdminBase(adminRender, msgRend, repoManager, store)
	adminGen := admin.NewAdminGeneral(&adminBase, repoManager)
	adminPosts := admin.NewAdminPost(&adminBase, repoManager)
	adminCat := admin.NewAdminCat(&adminBase, repoManager)
	adminUpload := admin.NewAdminUpload(&adminBase, repoManager)
	adminForgot := admin.NewAdminForgot(&adminBase, repoManager, emailMessenger)
	adminSeries := admin.NewAdminSeries(&adminBase, repoManager)

	frontBase := front.NewFrontBase(frontendRender, repoManager, config)
	frontPosts := front.NewFrontPosts(&frontBase)
	frontCategories := front.NewFrontCategories(&frontBase)

	app := milo.NewMiloApp(milo.SetPort(config.Port))

	adminGen.RegisterRoutes(app)
	adminPosts.RegisterRoutes(app)
	adminCat.RegisterRoutes(app)
	adminUpload.RegisterRoutes(app)
	adminForgot.RegisterRoutes(app)
	adminSeries.RegisterRoutes(app)
	frontBase.RegisterRoutes(app)
	frontPosts.RegisterRoutes(app)
	frontCategories.RegisterRoutes(app)

	app.RouteAssetStripPrefix("/admin", config.AdminDir)
	app.RouteAsset("/css", config.ThemeDir)
	app.RouteAsset("/js", config.ThemeDir)
	app.RouteAssetStripPrefix("/uploads", config.UploadDir)

	app.Run()
}
