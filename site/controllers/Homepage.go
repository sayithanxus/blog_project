package controllers

import (
	"blog_projesi/site/helpers"
	"blog_projesi/site/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"time"
)

type Homepage struct{}

func (homepage Homepage) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory": func(categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
		"getDate": func(t time.Time) string {
			return fmt.Sprintf("%02d.%02d.%d", t.Day(), int(t.Month()), t.Year())
		},
	}).ParseFiles(helpers.Include("homepage/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (homepage Homepage) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/detail")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get("slug = ?", params.ByName("slug"))
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) About(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/about")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["About"] = models.About{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}
func (homepage Homepage) Contact(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("homepage/Contact")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}

func (homepage Homepage) AddContact(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	namesurname := r.FormValue("name-surname")
	email := r.FormValue("email")
	message := r.FormValue("message")

	models.Contact{
		NameSurname: namesurname,
		Email:       email,
		Message:     message,
	}.Add()
	helpers.SetAlert(w, r, "Başarıyla Gönderildi")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
