# gator - Blog Aggregator

Learning Goals
Learn how to integrate a Go application with a PostgreSQL database.

Practice using your SQL skills to query and migrate a database (using sqlc and goose, two lightweight tools for typesafe SQL in Go).

Learn how to write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database


# psql basic commands

```bash
 psql postgres
 \c gator
 SELECT * FROM users;
```


# goose migrations

```bash
cd sql/schema
goose postgres <connection_string> up
goose postgres <connection_string> down
```

# go postgress driver

`go get github.com/lib/pq`
`import _ "github.com/lib/pq"`

"You have to import the driver, but you don't use it directly anywhere in your code. The underscore tells Go that you're importing it for its side effects, not because you need to use it."