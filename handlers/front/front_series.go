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

}

func (f FrontSeries) handleSingleSeries(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	series, err := f.rm.SeriesRepo.GetSeriesWithSlug(slug)
	if err != nil {
		f.RenderError(w, r, 500, err.Error())
		return
	}
	f.RenderJson(w, r, series)
}
