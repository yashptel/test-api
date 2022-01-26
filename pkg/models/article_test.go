package models

import (
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	articles = []Article{
		{
			ID:        1,
			Title:     "Article 1",
			SubTitle:  "Sub title 1",
			Content:   "Content 1",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Article 2",
			SubTitle:  "Sub title 2",
			Content:   "Content 2",
			CreatedAt: time.Now(),
		},
		{
			ID:        3,
			Title:     "Article 3",
			SubTitle:  "Sub title 3",
			Content:   "Content 3",
			CreatedAt: time.Now(),
		},
		{
			ID:        4,
			Title:     "Article 4",
			SubTitle:  "Sub title 4",
			Content:   "Content 4",
			CreatedAt: time.Now(),
		},
		{
			ID:        5,
			Title:     "Article 5",
			SubTitle:  "Sub title 5",
			Content:   "Content 5",
			CreatedAt: time.Now(),
		},
	}

	code := m.Run()
	os.Exit(code)
}

func Test_GetByID(t *testing.T) {
	article, err := NewArticleClient().GetByID(1)
	if err != nil {
		t.Error(err)
	}
	if article == nil {
		t.Error("Expected article to be not nil")
		return
	}

	if article.ID != 1 {
		t.Error("Expected ID to be 1")
	}
}

func Test_Create(t *testing.T) {
	article := &Article{
		Title:     "Article 6",
		SubTitle:  "Sub title 6",
		Content:   "Content 6",
		CreatedAt: time.Now(),
	}

	err := NewArticleClient().Create(article)
	if err != nil {
		t.Error(err)
	}

	if article.ID != 6 {
		t.Error("Expected ID to be 6")
	}
}

func Test_GetAll(t *testing.T) {
	res, err := NewArticleClient().GetAllWithPagination(1, 10)
	if err != nil {
		t.Error(err)
	}

	if len(res.Articles) != 5 {
		t.Error("Expected 5 articles but got", len(res.Articles))
	}

	if res.TotalPages != 1 {
		t.Error("Expected 1 articles")
	}

	if res.HasMore != false {
		t.Error("Expected no more articles")
	}

	if res.NextPage != 2 {
		t.Error("Expected next page to be 2")
	}
}

func Test_Search(t *testing.T) {
	res, err := NewArticleClient().Search("Article")
	if err != nil {
		t.Error(err)
	}

	if len(res) != 5 {
		t.Error("Expected 5 articles but got", len(res))
	}

	res, err = NewArticleClient().Search("Article 1")
	if err != nil {
		t.Error(err)
	}

	if len(res) != 1 {
		t.Error("Expected 1 article but got", len(res))
	}
}
