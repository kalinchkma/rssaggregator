package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nanashi10211/rssaggregator/internal/auth"
	"github.com/nanashi10211/rssaggregator/internal/database"
)

func (appCfg *appConfig) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Bad Request %s", err))
		return
	}
	// find user by email
	user, err := appCfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Bad Request %s", err))
		return
	}

	// compage user email
	if auth.ComparePassword(user.Password, params.Password) {
		cookie := &http.Cookie{
			Name:     "ApiKey",
			Value:    user.ApiKey,
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(w, cookie)
		respondWithJSON(w, http.StatusAccepted, databaseUserToUser(user))
		return
	}
}

func (appCfg *appConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Internal server error %s", err))
	}

	user, db_err := appCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  hashedPassword,
	})
	if db_err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn;t create a user: %s", db_err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (appCfg *appConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (appCfg *appConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Println(user.ID)
	posts, err := appCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get post: %v", err))
		return
	}
	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
