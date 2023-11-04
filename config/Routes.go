package config

import (
	admin "blog_projesi/admin/controllers"
	site "blog_projesi/site/controllers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMIN
	//Blog Posts
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	//Categories
	r.GET("/admin/kategoriler", admin.Categories{}.Index)
	r.POST("/admin/kategoriler/add", admin.Categories{}.Add)
	r.GET("/admin/kategoriler/delete/:id", admin.Categories{}.Delete)

	//Userops
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)
	r.GET("/admin/logout", admin.Userops{}.Logout)

	//AboutUs
	r.GET("/admin/aboutus", admin.AboutUs{}.Index)
	r.GET("/admin/aboutus/add", admin.AboutUs{}.NewItem)
	r.POST("/admin/aboutus/new", admin.AboutUs{}.Add)
	r.GET("/admin/aboutus/delete/:id", admin.AboutUs{}.Delete)
	r.GET("/admin/aboutus/edit/:id", admin.AboutUs{}.Edit)
	r.POST("/admin/aboutus/update/:id", admin.AboutUs{}.Update)
	//Contacts
	r.GET("/admin/contact", admin.Contact{}.Index)
	r.GET("/admin/contact/delete/:id", admin.Contact{}.Delete)
	//SITE
	//Homepage
	r.GET("/", site.Homepage{}.Index)
	r.GET("/yazilar/:slug", site.Homepage{}.Detail)
	r.GET("/hakkimizda", site.Homepage{}.About)
	r.GET("/iletisim", site.Homepage{}.Contact)
	r.POST("/site/contact/new", site.Homepage{}.AddContact)

	// SERVE FILES
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/assets/*filepath", http.Dir("site/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	return r
}
