package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"log"
	"net/http"
	"time"
)

type AdminPages struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminPages(base *AdminBase, rm usecases.RepoManager) AdminPages {
	ap := AdminPages{AdminBase: base, rm: rm}
	return ap
}

func (a AdminPages) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/pages", []string{"Get"}, a.authMid(a.handlePages))
	app.Route("/admin/pages/{id}/edit", []string{"Get"}, a.authMid(a.handlePageEdit))
	app.Route("/admin/pages/{id}/edit", []string{"Post"}, a.authMid(a.handlePageUpdate))
	app.Route("/admin/pages/start", []string{"Post"}, a.authMid(a.handlePageStart))
}

func (a AdminPages) handlePages(w http.ResponseWriter, r *http.Request) {
	data := a.setupActive("pages")
	pages, err := a.rm.PageRepo.FindAll()
	if err == nil && len(pages) > 0 {
		data["pages"] = pages
	}
	a.RenderTemplates(w, r, data, "base.tpl", "pages.tpl")
}

func (a AdminPages) handlePageEdit(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data := a.setupActive("pages")
	page, err := a.rm.PageRepo.FindById(id)
	if err == nil {
		data["page"] = page
	} else {
		log.Println("Page Load Error:", err)
	}
	a.RenderTemplates(w, r, data, "base.tpl", "edit_page.tpl")
}

func (a AdminPages) handlePageUpdate(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var page domain.Page
	err := dec.Decode(&page)
	if err != nil {
		a.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	upErr := a.rm.PageRepo.Update(&page)
	if upErr != nil {
		a.RenderError(w, r, http.StatusBadRequest, upErr.Error())
		return
	}
	a.RenderMessage(w, r, "Page Updated")
}

func (a AdminPages) handlePageStart(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form["title"][0]
	if title == "" {
		a.setErrorFlash(w, r, "Title required!")
		a.Redirect(w, r, "/admin/pages", http.StatusSeeOther)
		return
	}
	var page domain.Page
	page.Title = title
	page.Created = time.Now()

	page.GenerateSlug()

	err := a.rm.PageRepo.Store(&page)
	if err != nil {
		a.setErrorFlash(w, r, err.Error())
		a.Redirect(w, r, "/admin/pages", http.StatusSeeOther)
		return
	}

	a.setSuccessFlash(w, r, "Page created.")
	a.Redirect(w, r, fmt.Sprintf("/admin/pages/%d/edit", page.Id), http.StatusSeeOther)
}
