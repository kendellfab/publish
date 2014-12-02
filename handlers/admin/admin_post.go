package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/domain"
	"github.com/kendellfab/publish/usecases"
	"log"
	"net/http"
	"time"
)

type AdminPost struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminPost(base *AdminBase, rm usecases.RepoManager) AdminPost {
	post := AdminPost{AdminBase: base, rm: rm}
	return post
}

func (a AdminPost) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/posts", []string{"Get"}, a.authMid(a.handlePosts))
	app.Route("/admin/post/{id:[0-9]+}/edit", []string{"Get"}, a.authMid(a.handleEditPost))
	app.Route("/admin/post/{id:[0-9]+}/edit", []string{"Post"}, a.authMid(a.handleUpdatePost))
	app.Route("/admin/post/start", []string{"Post"}, a.authMid(a.handleStartPost))
}

func (a AdminPost) handlePosts(w http.ResponseWriter, r *http.Request) {
	posts, err := a.rm.PostRepo.FindAll()

	data := make(map[string]interface{})
	data["posts"] = posts
	if err != nil {
		data["error"] = err
	}

	a.RenderTemplates(w, r, data, "base.tpl", "posts.tpl")
}

func (a AdminPost) handleEditPost(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]
	post, err := a.rm.PostRepo.FindByIdString(idStr)

	data := make(map[string]interface{})
	data["post"] = post
	if err != nil {
		data["error"] = err
	}

	a.RenderTemplates(w, r, data, "base.tpl", "edit_post.tpl")
}

func (a AdminPost) handleUpdatePost(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var post domain.Post
	err := dec.Decode(&post)
	if err != nil {
		a.RenderError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	upErr := a.rm.PostRepo.Update(&post)
	if upErr != nil {
		a.RenderError(w, r, http.StatusBadRequest, upErr.Error())
		return
	}

	a.RenderMessage(w, r, "Post Updated")
}

func (a AdminPost) handleStartPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.Form["title"][0]

	if title == "" {
		a.setErrorFlash(w, r, "Title required.")
		a.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	post := &domain.Post{Title: title, Created: time.Now()}
	if ui, ok := context.GetOk(r, domain.CONTEXT_USER); ok {
		if usr, isUser := ui.(*domain.User); isUser {
			post.Author = *usr
			post.AuthorId = usr.Id
		} else {
			log.Println("Isn't a user.")
		}
	} else {
		log.Println("User not found.")
	}
	post.GenerateSlug()

	if strErr := a.rm.PostRepo.Store(post); strErr != nil {
		a.setErrorFlash(w, r, strErr.Error())
		log.Println("Start Post Error:", strErr.Error())
		a.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	a.Redirect(w, r, fmt.Sprintf("/admin/post/%d/edit", post.Id), http.StatusSeeOther)
}
