package handler

import (
	"log"
	"net/http"
	"text/template"
)

func (s *Server) getAbout(w http.ResponseWriter, req *http.Request) {
	tmp, _ := template.New("about.html").ParseFiles("./assets/templates/about.html")
	err := tmp.Execute(w, nil)
	if err != nil {
		log.Println("Error Executing Template : ", err)
		return
	}
}
