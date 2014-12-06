package front

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/usecases"
	"log"
	"net/http"
)

type FrontBase struct {
	milo.Renderer
	rm usecases.RepoManager
}

func NewFrontBase(rend milo.Renderer, rm usecases.RepoManager) FrontBase {
	base := FrontBase{Renderer: rend, rm: rm}
	return base
}

func (f FrontBase) RegisterRoutes(app *milo.Milo) {
	app.Route("/", nil, f.handleRoot)
}

func (f FrontBase) RenderTemplates(w http.ResponseWriter, r *http.Request, data map[string]interface{}, tpls ...string) {
	counts, countErr := f.rm.CategoryRepo.GetCategoryPostCount()
	if countErr == nil {
		data["Counts"] = counts
	}
	f.Renderer.RenderTemplates(w, r, data, tpls...)
}

func (f FrontBase) handleRoot(w http.ResponseWriter, r *http.Request) {
	posts, err := f.rm.PostRepo.FindPublished(0, 10)
	data := make(map[string]interface{})
	data["posts"] = posts
	if err != nil {
		log.Println(err)
		data["error"] = err
	}
	f.RenderTemplates(w, r, data, "index.html")
}
