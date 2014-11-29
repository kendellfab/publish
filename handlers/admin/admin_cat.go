package admin

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/usecases"
	"net/http"
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
}

func (a AdminCat) handlePosts(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "base.tpl", "cats.tpl")
}
