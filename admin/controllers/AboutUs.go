package controllers

import (
	"blog_projesi/admin/helpers"
	"blog_projesi/admin/models"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
)

type AboutUs struct{}

func (aboutus AboutUs) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("aboutus/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["AboutUs"] = models.About{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (aboutus AboutUs) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("aboutus/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["AboutUs"] = models.About{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (aboutus AboutUs) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	title := r.FormValue("aboutus-title")
	content := r.FormValue("aboutus-content")

	//Upload
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("aboutus-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file)
	// Upload End
	if err != nil {
		fmt.Println(err)
		return
	}
	models.About{
		Title:       title,
		Description: content,
		Picture_url: "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Kayıt Başarıyla Eklendi")
	http.Redirect(w, r, "/admin/aboutus", http.StatusSeeOther)
}

func (aboutus AboutUs) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	about := models.About{}.Get(params.ByName("id"))
	about.Delete()
	http.Redirect(w, r, "/admin/aboutus", http.StatusSeeOther)
}

func (aboutus AboutUs) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	view, err := template.ParseFiles(helpers.Include("aboutus/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["aboutus"] = models.About{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (aboutus AboutUs) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}
	about := models.About{}.Get(params.ByName("id"))
	title := r.FormValue("aboutus-title")
	content := r.FormValue("aboutus-content")
	is_selected := r.FormValue("is_selected")
	var picture_url string

	if is_selected == "1" {
		//Upload
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("aboutus-picture")
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, file)
		// Upload End
		picture_url = "uploads/" + header.Filename
		os.Remove(about.Picture_url)
	} else {
		picture_url = about.Picture_url
	}

	about.Update(models.About{
		Title:       title,
		Description: content,
		Picture_url: picture_url,
	})
	helpers.SetAlert(w, r, "Kayıt Başarıyla Güncellendi")
	http.Redirect(w, r, "/admin/aboutus", http.StatusSeeOther)
}
