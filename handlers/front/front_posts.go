package front

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"net/http"
	"time"
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
	year := mux.Vars(r)["year"]
	now := time.Now()

	series := make([]*domain.TimeSeries, 0)
	for i := int(now.Month()); i > 0; i-- {
		when, _ := time.Parse("2006-1", fmt.Sprintf("%s-%d", year, i))
		if posts, pErr := f.rm.PostRepo.FindByYearMonth(year, fmt.Sprintf("%d", i)); pErr == nil {
			series = append(series, &domain.TimeSeries{When: when, Posts: posts})
		}
	}
	data := make(map[string]interface{})
	data["series"] = series

	f.RenderTemplates(w, r, data, "time_series.html")
}

func (f FrontPosts) handleMonth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	when, _ := time.Parse("2006-01", fmt.Sprintf("%s-%s", year, month))

	posts, pErr := f.rm.PostRepo.FindByYearMonth(year, month)
	series := make([]*domain.TimeSeries, 0)
	series = append(series, &domain.TimeSeries{When: when, Posts: posts})

	data := make(map[string]interface{})
	data["series"] = series
	data["error"] = pErr

	f.RenderTemplates(w, r, data, "time_series.html")
}

func (f FrontPosts) handlePost(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	// view := &domain.View{Who: f.getIp(r), At: time.Now(), TargetType: domain.TypePost, Target: slug}
	// vErr := f.rm.ViewRepo.Store(view)
	// if vErr != nil {
	// 	log.Println(vErr)
	// }

	post, err := f.rm.PostRepo.FindBySlug(slug)
	data := make(map[string]interface{})
	if err == nil {
		data["Post"] = post
	} else {
		f.RenderError(w, r, 404, err.Error())
		return
	}

	if series, sErr := f.rm.SeriesRepo.GetSeries(fmt.Sprintf("%d", post.SeriesId)); sErr == nil {
		data["series"] = series
	}

	f.RenderTemplates(w, r, data, "post.html")
}
