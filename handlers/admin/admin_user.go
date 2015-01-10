package admin

import (
	"errors"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type AdminUser struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminuser(base *AdminBase, rm usecases.RepoManager) AdminUser {
	user := AdminUser{AdminBase: base, rm: rm}
	return user
}

func (a AdminUser) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/user/profile", []string{"Get"}, a.authMid(a.handleProfile))
	app.Route("/admin/user/profile/token/regen", []string{"Get"}, a.authMid(a.handleTokenRegen))
	app.Route("/admin/user/profile/update/password", []string{"Post"}, a.authMid(a.handleNewPassword))
	app.Route("/admin/user/profile/update/bio", []string{"Post"}, a.authMid(a.handleUpdateBio))

	app.Route("/admin/users", []string{"Get"}, a.authMid(a.handleUsers))
	app.Route("/admin/users/{id}/delete", []string{"Get"}, a.authMid(a.handleDelete))
	app.Route("/admin/users/add", []string{"Post"}, a.authMid(a.handleUsersAdd))
}

func (a AdminUser) handleProfile(w http.ResponseWriter, r *http.Request) {
	data := a.setupActive("users")
	if ui, ok := context.GetOk(r, domain.CONTEXT_USER); ok {
		if usr, isUser := ui.(*domain.User); isUser {
			data["user"] = usr
		} else {
			data["error"] = errors.New("Context: User not present")
		}
	}
	a.RenderTemplates(w, r, data, "base.tpl", "profile.tpl")
}

func (a AdminUser) handleTokenRegen(w http.ResponseWriter, r *http.Request) {
	if ui, ok := context.GetOk(r, domain.CONTEXT_USER); ok {
		if usr, isUser := ui.(*domain.User); isUser {
			usr.GenerateToken()
			a.rm.UserRepo.Update(usr)
			a.setSuccessFlash(w, r, "Token regenerated.")
		}
	}
	a.Redirect(w, r, "/admin/user/profile", http.StatusSeeOther)
}

func (a AdminUser) handleNewPassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	old := r.Form["old"][0]
	new1 := r.Form["new1"][0]
	new2 := r.Form["new2"][0]
	if new1 != new2 {
		a.setErrorFlash(w, r, "Your new passwords must match!")
	} else {
		if ui, ok := context.GetOk(r, domain.CONTEXT_USER); ok {
			if usr, isUser := ui.(*domain.User); isUser {
				if compErr := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(old)); compErr == nil {
					if bArr, bErr := bcrypt.GenerateFromPassword([]byte(new1), bcrypt.DefaultCost); bErr == nil {
						if upErr := a.rm.UserRepo.UpdatePassword(fmt.Sprintf("%d", usr.Id), string(bArr)); upErr == nil {
							a.setSuccessFlash(w, r, "Password updated.")
						} else {
							a.setErrorFlash(w, r, upErr.Error())
						}
					} else {
						a.setErrorFlash(w, r, bErr.Error())
					}
				} else {
					a.setErrorFlash(w, r, "Your old password didn't verify.")
				}
			}
		}
	}
	a.Redirect(w, r, "/admin/user/profile", http.StatusSeeOther)
}

func (a AdminUser) handleUpdateBio(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bio := r.Form["bio"][0]
	if ui, ok := context.GetOk(r, domain.CONTEXT_USER); ok {
		if usr, isUser := ui.(*domain.User); isUser {
			usr.Bio = bio
			if upErr := a.rm.UserRepo.Update(usr); upErr == nil {
				a.setSuccessFlash(w, r, "User bio updated.")
			} else {
				a.setErrorFlash(w, r, upErr.Error())
			}
		}
	}
	a.Redirect(w, r, "/admin/user/profile", http.StatusSeeOther)
}

func (a AdminUser) handleUsers(w http.ResponseWriter, r *http.Request) {
	data := a.setupActive("users")
	users, err := a.rm.UserRepo.GetAll()
	if err == nil {
		data["users"] = users
	} else {
		data["error"] = err
	}
	roles := make(map[int]string)
	roles[int(domain.Usr)] = domain.Usr.String()
	roles[int(domain.Author)] = domain.Author.String()
	roles[int(domain.Admin)] = domain.Admin.String()
	data["roles"] = roles
	a.RenderTemplates(w, r, data, "base.tpl", "users.tpl")
}

func (a AdminUser) handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	if ui, ok := context.GetOk(r, domain.CONTEXT_USER); ok {
		if usr, isUser := ui.(*domain.User); isUser {
			if usr.Id == int64(id) {
				a.setErrorFlash(w, r, "You can't delete yourself silly.")
			} else {
				delErr := a.rm.UserRepo.Delete(idStr)
				if delErr != nil {
					a.setErrorFlash(w, r, delErr.Error())
				} else {
					a.setSuccessFlash(w, r, "User deleted.")
				}
			}
		}
	} else {
		a.setErrorFlash(w, r, "Can't access logged in user.")
	}
	a.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

func (a AdminUser) handleUsersAdd(w http.ResponseWriter, r *http.Request) {
	// name, email, password, role
	r.ParseForm()
	name := r.Form["name"][0]
	email := r.Form["email"][0]
	password := r.Form["password"][0]
	rlStr := r.Form["role"][0]
	rlInt, _ := strconv.Atoi(rlStr)
	role := domain.Role(rlInt)
	fmt.Println("Role:", role)
	usr, usrErr := domain.NewUser(name, email, password, role)
	if usrErr != nil {
		a.setErrorFlash(w, r, usrErr.Error())
	} else {
		strErr := a.rm.UserRepo.Store(usr)
		if strErr != nil {
			a.setErrorFlash(w, r, strErr.Error())
		} else {
			a.setSuccessFlash(w, r, "User added!")
		}
	}

	a.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}
