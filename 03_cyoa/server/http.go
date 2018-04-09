package server

import (
	"html/template"
	"log"
	"net/http"
)

// Start initialize the HTTP Server
func Start(host, port string) {
	http.HandleFunc("/", handler)
	log.Println("[INIT] - HTTP Server")
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
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
