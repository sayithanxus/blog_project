package controllers

import (
	"blog_projesi/admin/helpers"
	"blog_projesi/admin/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type Contact struct{}

func (contact Contact) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("contact/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Contact"] = models.Contact{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (contact Contact) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	contacts := models.Contact{}.Get(params.ByName("id"))
	contacts.Delete()
	http.Redirect(w, r, "/admin/contact", http.StatusSeeOther)
}
