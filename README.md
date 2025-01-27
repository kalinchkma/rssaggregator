## Web Aggregator with `GO`

#### Tech used

* [GO]
* [goose]
* [sqlc]
* [postgresql]

#### Package used

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

#### Run project

Make copy of `.env.example` as `.env`

Install `goose` and `sqlc`

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


