# Go Articles API

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer.

After installing Go, run the following commands to start experiencing this starter kit:

```shell
# download the starter kit
git clone https://github.com/yashptel/test-api.git
cd test-api

# run the RESTful API server
go run main.go

```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

- `GET /healthz`: a healthcheck service provided for health checking purpose
- `POST api/articles`: create new article with body.
  ```json
  {
    "title": "Article 2",
    "sub_title": "Sub title 2",
    "content": "Content 2"
  }
  ```
- `GET /articles?page=1&limit=10`: returns a paginated list of the articles
- `GET /articles/:id`: returns the detailed information of an album
- `GET /articles/search?q=search_query`: search for article. It returns matching articles in fields: title, subtitle and content.

Try the URL `http://localhost:8080/healthz` in a browser, and you should see something like `"OK"` displayed.

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following
more complex scenarios:
