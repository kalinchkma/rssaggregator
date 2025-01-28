package main

import (
	"net/http"
	"text/template"

	"github.com/nanashi10211/rssaggregator/internal/database"
)

type HomePageData struct {
	PageTitle string
	Feeds     []database.Feed
	Posts     []database.Post
}

func (appConfig *appConfig) home(w http.ResponseWriter, r *http.Request, user database.User) {

	// get all feeds
	feeds, err := appConfig.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	posts, err := appConfig.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  1000,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	data := HomePageData{
		PageTitle: "Aggregator Home",
		Feeds:     feeds,
		Posts:     posts,
	}
	// parse template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, data)
}

type LoginPageData struct {
	PageTitle string
	Data      any
}

func (appConfig *appConfig) login(w http.ResponseWriter, r *http.Request) {
	data := LoginPageData{
		PageTitle: "Login ",
		Data:      "",
	}

	// parse template
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	tmpl.Execute(w, data)
}

type RegisterPageData struct {
	PageTitle string
	Data      any
}

func (appConfig *appConfig) register(w http.ResponseWriter, r *http.Request) {
	data := RegisterPageData{
		PageTitle: "Register",
		Data:      "",
	}

	// parse template
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	tmpl.Execute(w, data)
}
