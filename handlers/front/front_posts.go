package front

import (
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"net/http"
)

type FrontPosts struct {
	*FrontBase
}

func NewFrontPosts(base *FrontBase) FrontPosts {
	posts := FrontPosts{FrontBase: base}
	return posts
}

func (f FrontPosts) RegisterRoutes(app *milo.Milo) {
	app.Route("/{year:[0-9]+}", []string{"Get"}, f.handleYear)
	app.Route("/{year:[0-9]+}/{month:[0-9]+}", []string{"Get"}, f.handleMonth)
	app.Route("/{year:[0-9]+}/{month:[0-9]+}/{slug}", []string{"Get"}, f.handlePost)
}

func (f FrontPosts) handleYear(w http.ResponseWriter, r *http.Request) {

}

func (f FrontPosts) handleMonth(w http.ResponseWriter, r *http.Request) {

}

func (f FrontPosts) handlePost(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	post, err := f.rm.PostRepo.FindBySlug(slug)
	data := make(map[string]interface{})
	if err == nil {
		data["Post"] = post
	} else {
		data["PostError"] = err
	}
	f.RenderTemplates(w, r, data, "post.html")
}
