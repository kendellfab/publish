package front

import (
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"net/http"
)

type FrontSeries struct {
	*FrontBase
}

func NewFrontSeries(base *FrontBase) FrontSeries {
	series := FrontSeries{FrontBase: base}
	return series
}

func (f FrontSeries) RegisterRoutes(app *milo.Milo) {
	app.Route("/series", []string{"Get"}, f.handleSeries)
	app.Route("/series/{slug}", []string{"Get"}, f.handleSingleSeries)
}

func (f FrontSeries) handleSeries(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	series, sErr := f.rm.SeriesRepo.GetAll()
	data["series"] = series
	data["error"] = sErr
	f.RenderTemplates(w, r, data, "series_all.html")
}

func (f FrontSeries) handleSingleSeries(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	data := make(map[string]interface{})
	series, err := f.rm.SeriesRepo.GetSeriesWithSlug(slug)
	data["series"] = series
	data["error"] = err
	f.RenderTemplates(w, r, data, "series_one.html")
}
