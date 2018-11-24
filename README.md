# amzn-go-wish
A Go program to query the API provided by the amazon-wish-lister project that will be able to store the data to a CSV file.

## Setup

1. Build amazon-wish-lister docker image:
```
docker build -t amazon-wish-lister --file amazon-wish-lister_docker/Dockerfile amazon-wish-lister_docker/
```
2. Build scraper docker image:
```
docker build -t akumor/amzn-go-wish/scraper ./scraper/ 
```
3. Start containers with docker compose:
```
docker-compose -f ./docker-compose.yml
```

## Run

`go run main.go -host=localhost -port=8080 -id=XXXXXXXXXXXX -file=output.csv`

## Resouces
* amazon-wish-list https://github.com/doitlikejustin/amazon-wish-lister
