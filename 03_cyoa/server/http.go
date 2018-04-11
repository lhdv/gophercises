package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	cyoa "github.com/lhdv/gophercises/03_cyoa"
)

// Start initialize the HTTP Server
func Start(advFile, host, port string) {
	adv := cyoa.LoadAdventure(advFile)
	mux := defaultMux()

	handler := mapHandler(adv, mux)

	log.Println("[INIT] - HTTP Server")

	log.Fatal(http.ListenAndServe(host+":"+port, handler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootPage)
	return mux
}

func mapHandler(adventure cyoa.Adventure, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimLeft(r.URL.Path, "/")

		log.Println("[LOG] - mapHandler Call " + path)
		if dest, ok := adventure[path]; ok {
			buildPage(w, dest)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] - rootPage: Request")
	fmt.Fprintln(w, `<h1><a href="/intro">Start your own Adventure</a></h1>`)
}

func buildPage(w http.ResponseWriter, content cyoa.Content) *template.Template {
	t, err := template.ParseFiles("server/template.html")
	if err != nil {
		log.Fatal("[LOG] - Could not parse template file")
	}

	err = t.Execute(w, content)
	if err != nil {
		log.Fatal("[LOG] - Could not execute parsing on template")
	}

	return t
}
