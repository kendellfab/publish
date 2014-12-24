package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
	app.Route("/admin/series/{id}/edit", []string{"Post"}, a.authMid(a.handleSeriesUpdate))
	app.Route("/admin/series/start", []string{"Post"}, a.authMid(a.handleSeriesStart))
}

func (a AdminSeries) handleSeries(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	series, err := a.rm.SeriesRepo.GetAll()
	if err == nil {
		data["series"] = series
	}
	a.RenderTemplates(w, r, data, "base.tpl", "series.tpl")
}

func (a AdminSeries) handleSeriesEdit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data := make(map[string]interface{})
	if series, err := a.rm.SeriesRepo.GetSeries(id); err == nil {
		data["series"] = series
	}
	a.RenderTemplates(w, r, data, "base.tpl", "edit_series.tpl")
}

func (a AdminSeries) handleSeriesUpdate(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var series domain.Series
	err := dec.Decode(&series)
	if err != nil {
		a.RenderError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	upErr := a.rm.SeriesRepo.Update(&series)
	if upErr != nil {
		a.RenderError(w, r, http.StatusInternalServerError, upErr.Error())
		return
	}
	a.RenderMessage(w, r, "Series Updated")
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
