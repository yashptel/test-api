package controllers_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/yashptel/test-api/pkg/models"
)

func Test_ArticleController(t *testing.T) {

	t.Run("Test_Create", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, srv.URL+"/api/articles", nil)
		if err != nil {
			t.Error(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusBadRequest {
			t.Error("Expected status code to be 400")
		}

		payload, err := json.Marshal(models.Article{
			Title:    "test",
			SubTitle: "test sub title",
			Content:  "test content",
		})
		if err != nil {
			t.Error(err)
		}

		req, err = http.NewRequest(http.MethodPost, srv.URL+"/api/articles", bytes.NewBuffer(payload))
		if err != nil {
			t.Error(err)
		}
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Error("Expected status code to be 201")
		}
	})

	t.Run("Test_Get", func(t *testing.T) {

		payload, err := json.Marshal(models.Article{
			Title:    "test",
			SubTitle: "test sub title",
			Content:  "test content",
		})
		if err != nil {
			t.Error(err)
		}

		req, err := http.NewRequest(http.MethodPost, srv.URL+"/api/articles", bytes.NewBuffer(payload))
		if err != nil {
			t.Error(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Error("Expected status code to be 201")
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var article models.Article
		err = json.Unmarshal(bytes, &article)
		if err != nil {
			t.Error(err)
		}

		req, err = http.NewRequest(http.MethodGet, srv.URL+"/api/articles/"+strconv.Itoa(article.ID), nil)
		if err != nil {
			t.Error(err)
		}
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error("Expected status code to be 200")
		}

		bytes, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			t.Error(err)
		}

		var res models.Article
		err = json.Unmarshal(bytes, &res)
		if err != nil {
			t.Error(err)
		}
		if article.Title != res.Title {
			t.Error("Expected title to be test")
		}
		if article.SubTitle != res.SubTitle {
			t.Error("Expected sub title to be test sub title")
		}
		if article.Content != res.Content {
			t.Error("Expected content to be test content")
		}
	})

	t.Run("Test_GetAll", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, srv.URL+"/api/articles?page=1&limit=10", nil)
		if err != nil {
			t.Error(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error("Expected status code to be 200")
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var res models.ArticleList
		err = json.Unmarshal(bytes, &res)
		if err != nil {
			t.Error(err)
		}
		if len(res.Articles) < 2 {
			t.Error("Expected to have 1 article")
		}

	})

	t.Run("Test_Search", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, srv.URL+"/api/articles/search?q=test", nil)
		if err != nil {
			t.Error(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Error(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Error("Expected status code to be 200")
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		var res []models.Article
		err = json.Unmarshal(bytes, &res)
		if err != nil {
			t.Error(err)
		}
		if len(res) != 2 {
			t.Error("Expected to find 1 article but found", len(res))
		}
	})
}
