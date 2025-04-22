package handler

import (
	"go-project/storage/postgres"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/mux"
)

type Server struct {
	templates *template.Template
	store     *postgres.Storage
}

func NewServer(st *postgres.Storage) (*mux.Router, error) {
	s := &Server{
		store: st,
	}

	if err := s.ParseTemplates(); err != nil {
		return nil, err
	}

	req := mux.NewRouter()

	req.HandleFunc("/", s.getHome).Methods("GET")
	req.HandleFunc("/about", s.getAbout).Methods("GET")
	req.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
	return req, nil
}

func (s *Server) ParseTemplates() error {
	templates := template.New("templates").Funcs(template.FuncMap{
		"strrev": func(str string) string {
			n := len(str)
			runes := make([]rune, n)
			for _, rune := range str {
				n--
				runes[n] = rune
			}
			return string(runes[n:])
		},
	}).Funcs(sprig.FuncMap())

	tmpl, err := templates.ParseGlob("assets/templates/*.html")
	if err != nil {
		return err
	}
	s.templates = tmpl
	return nil
}
