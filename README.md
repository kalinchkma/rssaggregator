## Rss Feed Aggregator with `GO`

#### Tech stack

* [GO] 
* [goose]
* [sqlc]
* [postgresql]
* [godotenv]
* [chi]
* [cors]

[godotenv]: https://github.com/joho/godotenv
[chi]: https://github.com/go-chi/chi
[cors]: https://github.com/go-chi/cors
[sqlc]: https://sqlc.dev/
[postgresql]: https://www.postgresql.org/
[goose]: https://github.com/pressly/goose
[GO]: https://go.dev/

#### Features

- [x] Aggregate `xml` based RSS feed
- [x] Add RSS Feed
- [x] Scrape RSS Feed posts with certain amount of time frame
- [x] Register user
- [x] Authentication based on cookies
- [x] Follow unfollow rss feed 


#### Run project

Make copy of `.env.example` as `.env`

Install `goose` and `sqlc`

Create local copy for external package
```bash
go mod vendor
```
and run
```bash
go mod tidy
```

Go to `sql/schema` folder and run migrate

Migrate up

```bash
goose postgres postgres://<username>:<password>@<host>:<port>/<dbname> up
```
Migrate down

```bash
goose postgres postgres://<username>:<password>@<host>:<port>/<dbname> down
```

Start project

```bash
make start 
```
or 
```bash
go run ./*.go
```

Build database
```bash
make docker build
```
ok
```bash
docker compose build
```

Up database
```bash
make docker-up
```
or
```bash
docker compose up -d
```


