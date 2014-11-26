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
}

func (a AdminGeneral) handleAdmin(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "base.tpl", "index.tpl")
}
