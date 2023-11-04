package main

import (
	admin_models "blog_projesi/admin/models"
	"blog_projesi/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	admin_models.About{}.Migrate()
	//admin_models.Contact{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
