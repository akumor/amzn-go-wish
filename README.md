# amzn-go-wish
Repo for working with Amazon wishlist using Go

**not yet working**

## Initial Idea

Write a Go program to query the API provided by the amazon-wishlist-project that will be able to store the data to multiple backends starting with PostgreSQL.

## Setup

1. Build amazon-wish-lister docker image:
```
docker build -t amazon-wish-lister --file amazon-wish-lister_docker/amazon-wish-lister-Dockerfile amazon-wish-lister_docker/
```
2. Build scraper docker image:
```
docker build -t akumor/amzn-go-wish/scraper ./scraper/ 
```
3. Start containers with docker compose:
```
docker-compose -f ./docker-compose.yml
```

## Resouces
* amazon-wish-list https://github.com/doitlikejustin/amazon-wish-lister
