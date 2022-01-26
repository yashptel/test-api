package controllers

import (
	"net/http"

	"github.com/yashptel/test-api/pkg/utils"
)

const prefix = ""

func NewRouter() *http.ServeMux {

	mux := http.NewServeMux()

	SetupHealthCheck(mux)
	SetupArticleRoutes(mux)
	return mux
}

func SetupHealthCheck(mux *http.ServeMux) {
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		utils.JSONWrite(w, http.StatusOK, "OK")
	})
}

func SetupArticleRoutes(mux *http.ServeMux) {

	ac := NewArticleController()

	mux.HandleFunc(prefix+"/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ac.GetArticles(w, r)
		case "POST":
			ac.CreateArticle(w, r)
		default:
			utils.JSONError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})

	mux.HandleFunc(prefix+"/articles/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ac.GetArticle(w, r)
		default:
			utils.JSONError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})

	mux.HandleFunc(prefix+"/articles/search", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ac.SearchArticles(w, r)
		default:
			utils.JSONError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		}
	})

}
