package front

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type FrontBase struct {
	*milo.Renderer
	rm        usecases.RepoManager
	config    domain.Config
	pageCount int
}

func NewFrontBase(rend *milo.Renderer, rm usecases.RepoManager, c domain.Config) FrontBase {
	base := FrontBase{Renderer: rend, rm: rm, config: c}
	base.pageCount = 10
	if c.AppConfig.PerPage != 0 {
		base.pageCount = c.AppConfig.PerPage
	}
	return base
}

func (f FrontBase) RegisterRoutes(app *milo.Milo) {
	app.Route("/", nil, f.handleRoot)
	app.PathPrefix("/", nil, f.handlePages)
}

func (f FrontBase) RenderTemplates(w http.ResponseWriter, r *http.Request, data map[string]interface{}, tpls ...string) {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["payload"] = f.rm.PayloadRepo.GetPayload()
	data["Now"] = time.Now()
	f.Renderer.RenderTemplates(w, r, data, tpls...)
}

func (f FrontBase) handleRoot(w http.ResponseWriter, r *http.Request) {
	total, _ := f.rm.PostRepo.PublishedCount()
	paginator := GetPagination(r, total, f.pageCount)
	offset := paginator.Offset * paginator.Count
	count := paginator.Count
	posts, err := f.rm.PostRepo.FindPublished(offset, count)

	data := make(map[string]interface{})
	data["pagination"] = paginator
	data["posts"] = posts
	if err != nil {
		log.Println(err)
		data["error"] = err
	}
	f.RenderTemplates(w, r, data, "index.html")
}

func (f FrontBase) handlePages(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(r.RequestURI, "/")[1:]
	log.Println("Len:", len(path), "Path:", path)

	data := make(map[string]interface{})
	page, err := f.rm.PageRepo.FindBySlug(path[0])
	if err == nil {
		data["page"] = page
	} else {
		log.Println(err)
		data["error"] = err
	}

	f.RenderTemplates(w, r, data, "page.html")

}

func (f FrontBase) getIp(r *http.Request) string {
	if ipProxy := r.Header.Get("X-Real-IP"); len(ipProxy) > 0 {
		return ipProxy
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
