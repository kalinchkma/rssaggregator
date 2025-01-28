package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/nanashi10211/rssaggregator/internal/database"
	"github.com/nanashi10211/rssaggregator/internal/env"

	_ "github.com/lib/pq"
)

type appConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	PORT, err := env.GetString("PORT")
	if err != nil {
		log.Fatal(err)
	}

	dbURL, err := env.GetString("DB_URL")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	// connection conversion
	db := database.New(conn)

	app := appConfig{
		DB: db,
	}

	go strartScraping(
		db, 10, time.Minute,
	)

	// router that handle request
	router := chi.NewRouter()

	// cors middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// api routes
	apiRouterV1 := chi.NewRouter()
	apiRouterV1.Get("/healthz", handlerReadiness)
	apiRouterV1.Get("/err", handlerErr)

	apiRouterV1.Post("/users", app.handlerCreateUser)
	apiRouterV1.Get("/users", app.middlewareAuth(app.handlerGetUser))

	apiRouterV1.Post("/feeds", app.middlewareAuth(app.handlerCreateFeed))
	apiRouterV1.Get("/feeds", app.handlerGetFeeds)

	apiRouterV1.Get("/posts", app.middlewareAuth(app.handlerGetPostsForUser))

	apiRouterV1.Post("/feed_follows", app.middlewareAuth(app.handlerCreateFeedFollow))
	apiRouterV1.Get("/feed_follows", app.middlewareAuth(app.handlerGetFeedFollows))
	apiRouterV1.Delete("/feed_follows/{feedFollowID}", app.middlewareAuth(app.handlerDeleteFeedFollow))

	router.Mount("/api/v1", apiRouterV1)

	// Web routes
	webRouter := chi.NewRouter()

	// landing page
	webRouter.Get("/", app.webMiddlewareAuth(app.home))

	// Auth routes
	webRouter.Get("/login", app.login)
	webRouter.Get("/register", app.register)

	router.Mount("/", webRouter)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + PORT,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server strarting on port %v", PORT)

	serr := srv.ListenAndServe()

	if serr != nil {
		log.Fatal(serr)
	}

}
