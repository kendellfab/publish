package admin

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"time"
)

type AdminForgot struct {
	*AdminBase
	rm usecases.RepoManager
	em usecases.Emailer
}

func NewAdminForgot(base *AdminBase, rm usecases.RepoManager, em usecases.Emailer) AdminForgot {
	af := AdminForgot{AdminBase: base, rm: rm, em: em}
	return af
}

func (a AdminForgot) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/forgot", []string{"Get"}, a.handleForgot)
	app.Route("/admin/forgot", []string{"Post"}, a.handleSubmitForgot)
	app.Route("/admin/redeem/{token}", []string{"Get"}, a.handleRedeem)
	app.Route("/admin/redeem", []string{"Post"}, a.handleSubmitRedeem)
}

func (a AdminForgot) handleForgot(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "no_auth_base.tpl", "forgot.tpl")
}

func (a AdminForgot) handleSubmitForgot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["email"][0]
	user, err := a.rm.UserRepo.FindByEmail(email)
	if err != nil {
		a.setErrorFlash(w, r, err.Error())
		a.Redirect(w, r, "/admin/forgot", http.StatusSeeOther)
		return
	}

	reset := &domain.Reset{}
	reset.UserId = user.Id
	reset.Created = time.Now()
	dur, _ := time.ParseDuration("24h")
	reset.Expires = reset.Created.Add(dur)

	k := make([]byte, 25)
	io.ReadFull(rand.Reader, k)
	reset.Token = base64.StdEncoding.EncodeToString(k)

	strErr := a.rm.ResetRepo.Store(reset)
	if strErr != nil {
		a.setErrorFlash(w, r, strErr.Error())
		a.Redirect(w, r, "/admin/forgot", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["user"] = user
	data["request"] = r
	data["reset"] = reset

	msg, msgErr := a.MsgRender.RenderHtml(data, "reset_password.tpl")
	if msgErr != nil {
		a.setErrorFlash(w, r, msgErr.Error())
		a.Redirect(w, r, "/admin/forgot", http.StatusSeeOther)
		return
	}

	sendErr := a.em.SendMessage(user.Email, "Password Reset", msg)
	if sendErr != nil {
		a.setErrorFlash(w, r, sendErr.Error())
		a.Redirect(w, r, "/admin/forgot", http.StatusSeeOther)
		return
	}
	a.setSuccessFlash(w, r, "Password reset sent.")
	a.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (a AdminForgot) handleRedeem(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	data := make(map[string]interface{})
	reset, err := a.rm.ResetRepo.FindByToken(token)
	if err != nil {
		a.setErrorFlash(w, r, err.Error())
		a.Redirect(w, r, "/admin/forgot", http.StatusSeeOther)
		return
	}
	now := time.Now()
	if now.After(reset.Expires) {
		a.setErrorFlash(w, r, "Token has expired")
		a.Redirect(w, r, "/admin/forgot", http.StatusSeeOther)
		return
	}
	data["reset"] = reset
	user, uErr := a.rm.UserRepo.FindById(fmt.Sprintf("%d", reset.UserId))
	if uErr != nil {
		a.setErrorFlash(w, r, uErr.Error())
		a.Redirect(w, r, "/admin/forgot", http.StatusSeeOther)
		return
	}
	data["user"] = user

	a.RenderTemplates(w, r, data, "no_auth_base.tpl", "redeem.tpl")
}

func (a AdminForgot) handleSubmitRedeem(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userId := r.Form["user_id"][0]
	password := r.Form["password"][0]
	token := r.Form["token"][0]

	if password == "" {
		a.setErrorFlash(w, r, "Password must be set.")
		a.Redirect(w, r, "/admin/redeem/"+token, http.StatusSeeOther)
		return
	}

	bArr, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		a.setErrorFlash(w, r, err.Error())
		a.Redirect(w, r, "/admin/redeem/"+token, http.StatusSeeOther)
		return
	}

	upErr := a.rm.UserRepo.UpdatePassword(userId, string(bArr))
	if upErr != nil {
		a.setErrorFlash(w, r, upErr.Error())
		a.Redirect(w, r, "/admin/redeem/"+token, http.StatusSeeOther)
		return
	}

	a.setSuccessFlash(w, r, "Password reset")
	a.Redirect(w, r, "/login", http.StatusSeeOther)

}
