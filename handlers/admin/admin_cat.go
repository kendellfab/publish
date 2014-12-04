package admin

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"net/http"
	"time"
)

type AdminCat struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminCat(base *AdminBase, rm usecases.RepoManager) AdminCat {
	post := AdminCat{AdminBase: base, rm: rm}
	return post
}

func (a AdminCat) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/cats", []string{"Get"}, a.authMid(a.handlePosts))
	app.Route("/admin/cat/save", []string{"Post"}, a.authMid(a.handleSavePost))
}

func (a AdminCat) handlePosts(w http.ResponseWriter, r *http.Request) {
	cats, err := a.rm.CategoryRepo.GetAll()
	data := make(map[string]interface{})
	data["cats"] = cats
	data["error"] = err
	a.RenderTemplates(w, r, data, "base.tpl", "cats.tpl")
}

func (a AdminCat) handleSavePost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	title := r.Form["title"][0]

	var cat domain.Category

	cat.Title = title
	cat.Created = time.Now()

	if err := a.rm.CategoryRepo.Store(&cat); err != nil {
		a.setErrorFlash(w, r, err.Error())
	} else {
		a.setSuccessFlash(w, r, "Category saved.")
	}

	a.Redirect(w, r, "/admin/cats", http.StatusSeeOther)
}
