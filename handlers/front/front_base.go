package front

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/usecases"
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
	f.Renderer.RenderTemplates(w, r, data, tpls...)
}

func (f FrontBase) handleRoot(w http.ResponseWriter, r *http.Request) {
	f.RenderTemplates(w, r, nil, "base.tpl", "index.tpl")
}
