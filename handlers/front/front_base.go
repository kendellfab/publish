package front

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type FrontBase struct {
	milo.Renderer
	rm      usecases.RepoManager
	config  domain.Config
	perPage int
}

func NewFrontBase(rend milo.Renderer, rm usecases.RepoManager, c domain.Config) FrontBase {
	base := FrontBase{Renderer: rend, rm: rm, config: c}
	base.perPage = c.PerPage
	return base
}

func (f FrontBase) RegisterRoutes(app *milo.Milo) {
	app.Route("/", nil, f.handleRoot)
}

func (f FrontBase) RenderTemplates(w http.ResponseWriter, r *http.Request, data map[string]interface{}, tpls ...string) {
	if data == nil {
		data = make(map[string]interface{})
	}
	payStart := time.Now()
	data["payload"] = f.rm.PayloadRepo.GetPayload()
	log.Println(time.Since(payStart))
	data["Now"] = time.Now()
	f.Renderer.RenderTemplates(w, r, data, tpls...)
}

func (f FrontBase) getPagination(r *http.Request, total int) domain.Pagination {
	paginator := domain.Pagination{}
	paginator.Count = f.perPage

	r.ParseForm()
	if pStr, ok := r.Form["page"]; ok {
		if pg, pgErr := strconv.Atoi(pStr[0]); pgErr == nil {
			paginator.Offset = pg
		}
	}

	if paginator.Offset > 1 {
		paginator.HasNewer = true
		paginator.NewerIndex = paginator.Offset - 1
	}

	if paginator.Offset*paginator.Count < total {
		paginator.HasOlder = true
		paginator.OlderIndex = paginator.Offset + 1
		if paginator.OlderIndex == 1 {
			paginator.OlderIndex += 1
		}
	}

	paginator.Total = total

	return paginator
}

func (f FrontBase) handleRoot(w http.ResponseWriter, r *http.Request) {
	total, _ := f.rm.PostRepo.PublishedCount()
	paginator := f.getPagination(r, total)
	posts, err := f.rm.PostRepo.FindPublished((paginator.Offset-1)*paginator.Count, paginator.Count)
	data := make(map[string]interface{})
	data["pagination"] = paginator
	data["posts"] = posts
	if err != nil {
		log.Println(err)
		data["error"] = err
	}
	f.RenderTemplates(w, r, data, "index.html")
}

func (f FrontBase) getIp(r *http.Request) string {
	if ipProxy := r.Header.Get("X-Real-IP"); len(ipProxy) > 0 {
		return ipProxy
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
