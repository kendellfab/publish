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
	app.Route("/admin/post/{id:[0-9]+}/publish", []string{"Post"}, a.authMid(a.handlePublishPost))
	app.Route("/admin/post/{id:[0-9]+}/delete", []string{"Get"}, a.authMid(a.handleDeletePost))
	app.Route("/admin/post/start", []string{"Post"}, a.authMid(a.handleStartPost))
}

func (a AdminPost) handlePosts(w http.ResponseWriter, r *http.Request) {
	posts, err := a.rm.PostRepo.FindAll()

	// data := make(map[string]interface{})
	data := a.setupActive("post")
	data["posts"] = posts
	if err != nil {
		data["error"] = err
	}

	a.RenderTemplates(w, r, data, "base.tpl", "posts.tpl")
}

func (a AdminPost) handleEditPost(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]
	post, err := a.rm.PostRepo.FindByIdString(idStr)

	// data := make(map[string]interface{})
	data := a.setupActive("post")
	data["post"] = post
	if err != nil {
		data["error"] = err
	}

	cats, cErr := a.rm.CategoryRepo.GetAll()
	if cErr == nil {
		data["cats"] = cats
	} else {
		log.Println(cErr)
	}

	if series, sErr := a.rm.SeriesRepo.GetAll(); sErr == nil {
		data["series"] = series
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

func (a AdminPost) handlePublishPost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	dec := json.NewDecoder(r.Body)
	decErr := dec.Decode(&post)
	if decErr != nil {
		a.RenderError(w, r, 500, decErr.Error())
		return
	}

	if post.Published {
		pubErr := a.rm.PostRepo.Publish(post.Id)
		if pubErr != nil {
			a.RenderError(w, r, 500, pubErr.Error())
			return
		}
		a.RenderMessage(w, r, "Post published")
		return
	}

	unPubErr := a.rm.PostRepo.UnPublish(post.Id)
	if unPubErr != nil {
		a.RenderError(w, r, 500, unPubErr.Error())
		return
	}
	a.RenderMessage(w, r, "Post unpublished.")
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

func (a AdminPost) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := a.rm.PostRepo.Delete(id)
	if err != nil {
		a.setErrorFlash(w, r, err.Error())
	} else {
		a.setSuccessFlash(w, r, "Post Deleted")
	}
	a.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}
