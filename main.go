package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/ChabuChabu/lenslocked/controllers"
	"github.com/ChabuChabu/lenslocked/views"
	"github.com/go-chi/chi/v5"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "aboutUs.gohtml")))
	r.Get("/about", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port:3000....")
	http.ListenAndServe(":3000", r)
}
