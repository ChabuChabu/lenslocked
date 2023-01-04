package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
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

// func pathhandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		//handle the page not found error
// 		http.Error(w, "page not found 3", http.StatusNotFound)

// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		//handle the page not found error
		http.Error(w, "page not found 3", http.StatusNotFound)

	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on port:3000....")
	http.ListenAndServe(":3000", router)
}
