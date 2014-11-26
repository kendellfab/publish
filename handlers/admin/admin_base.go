package admin

import (
	"github.com/gorilla/sessions"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"net/http"
)

type AdminBase struct {
	milo.Renderer
	rm    usecases.RepoManager
	store sessions.Store
}

func NewAdminBase(rend milo.Renderer, rm usecases.RepoManager, store sessions.Store) AdminBase {
	base := AdminBase{Renderer: rend, rm: rm, store: store}
	return base
}

func (a AdminBase) authMid(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

func (a AdminBase) RenderTemplates(w http.ResponseWriter, r *http.Request, data map[string]interface{}, tpls ...string) {
	if data == nil {
		data = make(map[string]interface{})
	}

	if sess, sessErr := a.store.Get(r, domain.SESS_FLASH); sessErr == nil {
		data[domain.FLASH_ERROR] = sess.Flashes(domain.FLASH_ERROR)
		data[domain.FLASH_SUCCESS] = sess.Flashes(domain.FLASH_SUCCESS)
		sess.Options.MaxAge = -1
		sess.Save(r, w)
	}

	a.Renderer.RenderTemplates(w, r, data, tpls...)
}

func (a AdminBase) setErrorFlash(w http.ResponseWriter, r *http.Request, message string) {
	a.setFlashMessage(w, r, domain.FLASH_ERROR, message)
}

func (a AdminBase) setSuccessFlash(w http.ResponseWriter, r *http.Request, message string) {
	a.setFlashMessage(w, r, domain.FLASH_SUCCESS, message)
}

func (a AdminBase) setFlashMessage(w http.ResponseWriter, r *http.Request, key, message string) {
	if sess, sessErr := a.store.Get(r, domain.SESS_FLASH); sessErr == nil {
		sess.AddFlash(message, key)
		sess.Save(r, w)
	}
}
