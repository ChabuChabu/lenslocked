package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	bio := `&lt;script&gt;alert(&quot;Hi!&quot;);&lt;/script&gt;`
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1><p>User's bio: "+bio+"</p>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1> Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:chabuchabu32@gmail.com\">chabucabu32@gmail.com</a>.")

}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1> Frequently asked questions</h1>
	 <ul>
	   <li>is there a free version for this ? Yes! we offer a trial of 30 days for any paid plan.</li>
	 </ul>
	 <ul>
	   <li> Are there any prerequisites ? Absolutely not! we explain everything from ground up.</li>
     </ul>
     `)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port:3000....")
	http.ListenAndServe(":3000", r)
}
