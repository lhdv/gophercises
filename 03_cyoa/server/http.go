package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	cyoa "github.com/lhdv/gophercises/03_cyoa"
)

// Start initialize the HTTP Server
func Start(host, port string) {
	mux := defaultMux()

	log.Println("[INIT] - HTTP Server")

	log.Fatal(http.ListenAndServe(host+":"+port, mux))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootPage)
	mux.HandleFunc("/start", startPage)
	return mux
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] - rootPage: Request")
	fmt.Fprintln(w, `<h1><a href="/start">Start your own Adventure</a></h1>`)
}

func startPage(w http.ResponseWriter, r *http.Request) {

	adv := cyoa.LoadAdventure("gopher.json")

	t, err := template.ParseFiles("server/template.html")
	if err != nil {
		log.Fatal("[LOG] - Could not parse template file")
	}
	// _ = t
	t.Execute(w, adv["intro"])
	log.Printf("[LOG] - %+v\n", adv["intro"])
}

func advHandler(fallback http.HandlerFunc) http.HandlerFunc {

	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		fallback.ServeHTTP(w, r)
	}

	return handleFunc
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] - Request")

	p := struct {
		Title string
		Story []string
	}{
		Title: "Test",
		Story: []string{"Story line 001", "Story line 002", "Story line 003"},
	}

	t, err := template.ParseFiles("server/template.html")
	if err != nil {
		log.Fatal("[LOG] - Could not parse template file")
	}
	// _ = t
	t.Execute(w, p)
	log.Printf("[LOG] - %+v\n", p)
}
