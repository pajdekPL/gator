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

# sqlc code generation 

generate code based on the files in the `sql/queries` dir

```bash 
sqlc generate
```
# go postgress driver

```bash
go get github.com/lib/pq
import _ "github.com/lib/pq"
```

"You have to import the driver, but you don't use it directly anywhere in your code. The underscore tells Go that you're importing it for its side effects, not because you need to use it."

# Installation

To install the gator CLI, make sure you have Go installed on your system, then run:

```bash
go install github.com/pajdekpl/gator@latest
```

This will download and install the latest version of gator. The binary will be installed to your `$GOPATH/bin` directory. Make sure this directory is in your system's PATH to use the `gator` command from anywhere.

# Usage Examples

## User Management

Register a new user:
```bash
gator register johndoe
```

Login as an existing user:
```bash
gator login johndoe
```

List all users:
```bash
gator users
```

## Feed Management

Add a new RSS feed:
```bash
gator add-feed "Golang Blog" https://go.dev/blog/feed.atom
```

List all feeds:
```bash
gator feeds
```

## Following Feeds

Follow a feed:
```bash
gator follow "Golang Blog"
```

Unfollow a feed:
```bash
gator unfollow "Golang Blog"
```

List followed feeds:
```bash
gator following
```

## Reading Posts

Browse posts from followed feeds:
```bash
gator browse
```

## Aggregation

Manually trigger feed aggregation:
```bash
gator agg
```

The service will also automatically fetch new posts from RSS feeds periodically.