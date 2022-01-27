package models

import (
	"strings"
	"sync"
	"time"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	SubTitle  string    `json:"sub_title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleList struct {
	Articles   []Article `json:"articles"`
	HasMore    bool      `json:"has_more"`
	NextPage   int       `json:"next_page"`
	TotalPages int       `json:"total_pages"`
}

type ArticleRepository interface {
	GetAllWithPagination(page int, limit int) (ArticleList, error)
	GetByID(id int) (*Article, error)
	Create(article *Article) error
	Search(query string) ([]Article, error)
}

type ArticleClient struct {
}

func NewArticleClient() ArticleRepository {
	return &ArticleClient{}
}

var mu sync.Mutex
var articles = []Article{
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

func (c *ArticleClient) GetByID(id int) (*Article, error) {
	for _, article := range articles {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, nil
}

func (c *ArticleClient) Create(article *Article) error {
	article.ID = len(articles) + 1
	article.CreatedAt = time.Now()
	mu.Lock()
	defer mu.Unlock()
	articles = append(articles, *article)
	return nil
}

func (c *ArticleClient) Search(query string) ([]Article, error) {

	res := make([]Article, 0)
	query = strings.ToLower(query)

	for _, article := range articles {
		if strings.Contains(strings.ToLower(article.Title), query) ||
			strings.Contains(strings.ToLower(article.SubTitle), query) ||
			strings.Contains(strings.ToLower(article.Content), query) {
			res = append(res, article)
		}
	}
	return res, nil
}

func (c *ArticleClient) GetAllWithPagination(page int, limit int) (ArticleList, error) {
	var res ArticleList
	start := (page - 1) * limit
	end := start + limit

	if start > len(articles) {
		start = len(articles)
	}

	if end > len(articles) {
		end = len(articles)
	}

	res.Articles = articles[start:end]
	res.HasMore = end < len(articles)
	res.NextPage = page + 1

	if len(articles)%limit == 0 {
		res.TotalPages = len(articles) / limit
	} else {
		res.TotalPages = len(articles)/limit + 1
	}

	return res, nil
}
