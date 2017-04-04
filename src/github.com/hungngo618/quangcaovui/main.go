package main

import (
	"html/template"
	"log"
	"net/http"
)
const staticPath = "src/github.com/hungngo618/quangcaovui/ui/resources"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(staticPath+"/index.html")
	if err != nil {
		log.Fatal("fail to render template: ", err.Error())
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("failed to exe: ", err.Error())
	}
}


func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir(staticPath))))
	http.HandleFunc("/", IndexHandler)
	log.Println("Server is running at http://localhost:1203")
	log.Println(http.ListenAndServe(":1203", nil))
}
