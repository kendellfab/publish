package admin

import (
	"fmt"
	"github.com/kendellfab/milo"
	"github.com/kendellfab/publish/usecases"
	"net/http"
)

const K24 = (1 << 20) * 24

type AdminUpload struct {
	*AdminBase
	rm usecases.RepoManager
}

func NewAdminUpload(base *AdminBase, rm usecases.RepoManager) AdminUpload {
	up := AdminUpload{AdminBase: base, rm: rm}
	return up
}

func (a AdminUpload) RegisterRoutes(app *milo.Milo) {
	app.Route("/admin/uploads", []string{"Get"}, a.authMid(a.handleUploads))
	app.Route("/admin/upload", []string{"Post"}, a.authMid(a.handleDoUpload))
}

func (a AdminUpload) handleUploads(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplates(w, r, nil, "base.tpl", "uploads.tpl")
}

func (a AdminUpload) handleDoUpload(w http.ResponseWriter, r *http.Request) {
	if pErr := r.ParseMultipartForm(K24); pErr != nil {
		a.setErrorFlash(w, r, pErr.Error())
	} else {
		for _, fHeader := range r.MultipartForm.File {
			for _, header := range fHeader {
				if sErr := a.rm.UploadRepo.SaveMultipartFile(header); sErr == nil {
					a.setSuccessFlash(w, r, fmt.Sprintf("%s saved.", header.Filename))
				} else {
					a.setErrorFlash(w, r, fmt.Sprintf("%s not saved: %s", header.Filename, sErr.Error()))
				}
			}
		}
	}
	a.Redirect(w, r, "/admin/uploads", http.StatusSeeOther)
}
