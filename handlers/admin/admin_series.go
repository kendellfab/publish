package admin

import (
	"fmt"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"net/http"
	"time"
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
	app.Route("/admin/series/{id}/edit", []string{"Get"}, a.authMid(a.handleSeriesEdit))
	app.Route("/admin/series/start", []string{"Post"}, a.authMid(a.handleSeriesStart))
}

func (a AdminSeries) handleSeries(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "base.tpl", "series.tpl")
}

func (a AdminSeries) handleSeriesEdit(w http.ResponseWriter, r *http.Request) {

}

func (a AdminSeries) handleSeriesStart(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form["title"][0]

	if title == "" {
		a.setErrorFlash(w, r, "Title required")
		a.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	series := &domain.Series{Title: title, Created: time.Now()}
	series.GenerateSlug()

	if strErr := a.rm.SeriesRepo.Store(series); strErr != nil {
		a.setErrorFlash(w, r, strErr.Error())
		a.Redirect(w, r, "/admin/series", http.StatusSeeOther)
		return
	}

	a.Redirect(w, r, fmt.Sprintf("/admin/series/%d/edit", series.Id), http.StatusSeeOther)
}
