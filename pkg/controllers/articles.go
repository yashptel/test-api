package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yashptel/test-api/pkg/models"
	"github.com/yashptel/test-api/pkg/utils"
)

type ArticleController struct {
	ArticleModel models.ArticleRepository
}

func NewArticleController() *ArticleController {
	repo := models.NewArticleClient()
	return &ArticleController{
		ArticleModel: repo,
	}
}

func (a *ArticleController) GetArticles(w http.ResponseWriter, r *http.Request) {

	limit := utils.GetIntFromURL(r, "limit", 10)
	page := utils.GetIntFromURL(r, "page", 1)

	articles, err := a.ArticleModel.GetAllWithPagination(page, limit)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONWrite(w, http.StatusOK, articles)
}

func (a *ArticleController) GetArticle(w http.ResponseWriter, r *http.Request) {

	if arr := strings.Split(r.URL.Path, "/articles/"); len(arr) != 2 {
		utils.JSONError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	id, err := strconv.Atoi(strings.Split(r.URL.Path, "/articles/")[1])
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Invalid article id")
		return
	}

	article, err := a.ArticleModel.GetByID(id)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if article == nil {
		utils.JSONError(w, http.StatusNotFound, "Article not found")
		return
	}
	utils.JSONWrite(w, http.StatusOK, article)
}

func (a *ArticleController) CreateArticle(w http.ResponseWriter, r *http.Request) {

	article := models.Article{}
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = a.ArticleModel.Create(&article)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONWrite(w, http.StatusCreated, article)
}

func (a *ArticleController) SearchArticles(w http.ResponseWriter, r *http.Request) {

	searchQuery := r.URL.Query().Get("q")
	if searchQuery == "" {
		utils.JSONError(w, http.StatusBadRequest, "Bad Request")
		return
	}

	articles, err := a.ArticleModel.Search(searchQuery)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.JSONWrite(w, http.StatusOK, articles)
}
