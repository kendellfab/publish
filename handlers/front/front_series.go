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

	total, _ := f.rm.SeriesRepo.Count()
	paginator := GetPagination(r, total, f.pageCount)
	offset := paginator.Offset * paginator.Count

	series, err := f.rm.SeriesRepo.GetSeriesLimit(offset, paginator.Count)
	data["series"] = series
	data["pagination"] = paginator
	data["error"] = err
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
