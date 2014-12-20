package admin

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/usecases"
	"net/http"
)

type AdminSeries struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminSeries(base *AdminBase, rm usecases.RepoManager) AdminSeries {
	as := AdminSeries{AdminBase: base, rm: rm}
	return as
}

func (a AdminSeries) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/series", []string{"Get"}, a.authMid(a.handleSeries))
}

func (a AdminSeries) handleSeries(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "base.tpl", "series.tpl")
}
