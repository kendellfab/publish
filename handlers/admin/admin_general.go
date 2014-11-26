package admin

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/usecases"
	"net/http"
)

type AdminGeneral struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminGeneral(base *AdminBase, rm usecases.RepoManager) AdminGeneral {
	general := AdminGeneral{AdminBase: base, rm: rm}
	return general
}

func (a AdminGeneral) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin", []string{"Get"}, a.authMid(a.handleAdmin))
	app.Route("/setup", []string{"Get"}, a.handleSetup)
	app.Route("/login", []string{"Get"}, a.handleLogin)
	app.Route("/logout", []string{"Get"}, a.handleLogout)
}

func (a AdminGeneral) handleAdmin(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "base.tpl", "index.tpl")
}

func (a AdminGeneral) handleSetup(w http.ResponseWriter, r *http.Request) {
	// a.RenderMessage(w, r, "Setup")
	a.RenderTemplates(w, r, nil, "setup.tpl")
}

func (a AdminGeneral) handleLogin(w http.ResponseWriter, r *http.Request) {
	a.RenderMessage(w, r, "Login")
}

func (a AdminGeneral) handleLogout(w http.ResponseWriter, r *http.Request) {
	a.RenderMessage(w, r, "Logout")
}
