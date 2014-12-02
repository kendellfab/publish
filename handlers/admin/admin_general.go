package admin

import (
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AdminGeneral struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminGeneral(base *AdminBase, rm usecases.RepoManager) AdminGeneral {
	general := AdminGeneral{AdminBase: base, rm: rm}
	return general
}

func (a AdminGeneral) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin", []string{"Get"}, a.authMid(a.handleAdmin))
	app.Route("/setup", []string{"Get", "Post"}, a.handleSetup)
	app.Route("/login", []string{"Get", "Post"}, a.handleLogin)
	app.Route("/logout", []string{"Get"}, a.handleLogout)
}

func (a AdminGeneral) handleAdmin(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, a.setupActive("dashboard"), "base.tpl", "index.tpl")
}

func (a AdminGeneral) handleSetup(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	errs := make([]string, 0)

	if r.Method == "POST" {
		r.ParseForm()
		name := r.Form["name"][0]
		email := r.Form["email"][0]
		password := r.Form["password"][0]

		usr, usrErr := domain.NewAdminUser(name, email, password)
		if usrErr == nil {
			if strErr := a.rm.UserRepo.Store(usr); strErr == nil {
				a.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			} else {
				errs = append(errs, strErr.Error())
			}
		} else {
			errs = append(errs, usrErr.Error())
		}
	}
	if len(errs) > 0 {
		data["setup_error"] = errs
	}
	a.RenderTemplates(w, r, data, "setup.tpl")
}

func (a AdminGeneral) handleLogin(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	errs := make([]string, 0)

	if r.Method == "POST" {
		r.ParseForm()

		email := r.Form["email"][0]
		password := r.Form["password"][0]

		if usr, usrErr := a.rm.UserRepo.FindByEmail(email); usrErr == nil {
			if passwrdErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)); passwrdErr == nil {
				a.doLogin(usr, w, r)
				a.Redirect(w, r, "/admin", http.StatusSeeOther)
				return
			} else {
				errs = append(errs, passwrdErr.Error())
			}
		} else {
			errs = append(errs, usrErr.Error())
		}
	}

	if len(errs) > 0 {
		data["setup_error"] = errs
	}
	a.RenderTemplates(w, r, data, "login.tpl")
}

func (a AdminGeneral) handleLogout(w http.ResponseWriter, r *http.Request) {
	a.doLogout(w, r)
	a.setSuccessFlash(w, r, "Logged Out")
	a.Redirect(w, r, "/login", http.StatusSeeOther)
}
