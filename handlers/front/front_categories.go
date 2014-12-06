package front

import (
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"net/http"
)

type FrontCategories struct {
	*FrontBase
}

func NewFrontCategories(base *FrontBase) FrontCategories {
	posts := FrontCategories{FrontBase: base}
	return posts
}

func (f FrontCategories) RegisterRoutes(app *milo.Milo) {
	app.Route("/category", []string{"Get"}, f.handleCategories)
	app.Route("/category/{slug}", []string{"Get"}, f.handleCategory)
}

func (f FrontCategories) handleCategories(w http.ResponseWriter, r *http.Request) {

}

func (f FrontCategories) handleCategory(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	w.Write([]byte(slug))
}
