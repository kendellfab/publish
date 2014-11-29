package admin

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/usecases"
	"net/http"
)

type AdminPost struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminPost(base *AdminBase, rm usecases.RepoManager) AdminPost {
	post := AdminPost{AdminBase: base, rm: rm}
	return post
}

func (a AdminPost) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/posts", []string{"Get"}, a.authMid(a.handlePosts))
}

func (a AdminPost) handlePosts(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "base.tpl", "posts.tpl")
}
