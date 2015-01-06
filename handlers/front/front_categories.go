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
	f.RenderTemplates(w, r, nil, "categories.html")
}

func (f FrontCategories) handleCategory(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	data := make(map[string]interface{})
	cat, err := f.rm.CategoryRepo.FindBySlug(slug)
	if err == nil {
		total, _ := f.rm.PostRepo.PublishedCountCategory(cat.Id)
		paginator := GetPagination(r, total, f.pageCount)
		offset := paginator.Offset * paginator.Count
		count := paginator.Count

		data["category"] = cat
		data["pagination"] = paginator
		posts, pErr := f.rm.PostRepo.FindByCategory(cat, offset, count)
		if pErr == nil {
			data["posts"] = posts
		} else {
			data["error"] = pErr
		}
	} else {
		data["error"] = err
	}

	f.RenderTemplates(w, r, data, "category.html")
}
