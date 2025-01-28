package main

import (
	"fmt"
	"net/http"

	"github.com/nanashi10211/rssaggregator/internal/auth"
	"github.com/nanashi10211/rssaggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (appCfg *appConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKeyFromCookies(r.Cookies())
		if err != nil {
			// respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		user, err := appCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}

func (appConfig *appConfig) webMiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKeyFromCookies(r.Cookies())
		if err != nil {
			// respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user, err := appConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			// respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))

			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		handler(w, r, user)
	}
}
