package front

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"log"
	"net"
	"net/http"
	"time"
)

type FrontBase struct {
	milo.Renderer
	rm     usecases.RepoManager
	config domain.Config
}

func NewFrontBase(rend milo.Renderer, rm usecases.RepoManager, c domain.Config) FrontBase {
	base := FrontBase{Renderer: rend, rm: rm, config: c}
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

func (f FrontBase) getIp(r *http.Request) string {
	if ipProxy := r.Header.Get("X-Real-IP"); len(ipProxy) > 0 {
		return ipProxy
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
