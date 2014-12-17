package admin

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"net/http"
)

type AdminBase struct {
	*milo.Renderer
	*milo.MsgRender
	rm    usecases.RepoManager
	store sessions.Store
}

func NewAdminBase(rend *milo.Renderer, msgRend *milo.MsgRender, rm usecases.RepoManager, store sessions.Store) AdminBase {
	base := AdminBase{Renderer: rend, MsgRender: msgRend, rm: rm, store: store}
	return base
}

func (a AdminBase) authMid(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sess, sessErr := a.store.Get(r, domain.SESS_AUTH_KEY)
		if sessErr != nil {
			a.setErrorFlash(w, r, sessErr.Error())
			a.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		id, idOk := sess.Values[domain.SESS_USER_ID]

		if !idOk {
			a.setErrorFlash(w, r, r.RequestURI+" requires authentication.")
			a.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		usr, usrErr := a.rm.UserRepo.FindByIdInt(id.(int64))
		if usrErr != nil {
			a.setErrorFlash(w, r, usrErr.Error())
			a.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		context.Set(r, domain.CONTEXT_USER, usr)

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

	if usr, ok := context.GetOk(r, domain.CONTEXT_USER); ok {
		data[domain.CONTEXT_USER] = usr
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

func (a AdminBase) doLogin(user *domain.User, w http.ResponseWriter, r *http.Request) error {
	sess, sessErr := a.store.Get(r, domain.SESS_AUTH_KEY)
	if sessErr != nil {
		return sessErr
	}
	sess.Values[domain.SESS_USER_ID] = user.Id
	sess.Values[domain.SESS_LOGGED_IN] = true
	sess.Options.MaxAge = 60 * 60 * 2
	sess.Save(r, w)
	return nil
}

func (a AdminBase) doLogout(w http.ResponseWriter, r *http.Request) error {
	sess, sessErr := a.store.Get(r, domain.SESS_AUTH_KEY)
	if sessErr != nil {
		return sessErr
	}
	sess.Options.MaxAge = -1
	sess.Save(r, w)
	return nil
}

func (a AdminBase) setupActive(active string) map[string]interface{} {
	res := make(map[string]interface{})
	res[domain.ACTIVE] = active
	return res
}
