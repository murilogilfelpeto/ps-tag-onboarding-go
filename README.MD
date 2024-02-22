# ps-tag-onboarding-go

## The project
This is the Go version of the [onboarding exercise](https://github.com/wexinc/ps-tag-onboarding) for the Tag Team.

## Frameworks, libs and technologies
- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [Galidator](https://github.com/golodash/galidator)
- [Swagger](https://github.com/swaggo/gin-swagger)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/)
- [Viper](https://github.com/spf13/viper)
- [Logrus](https://github.com/Sirupsen/logrus)

## Technical Proposal
- [Swagger](https://docs.google.com/document/d/1K1WaCZhHi35BbvQ7UAGWOlJeGBoIlLVXpW7EQ8ryFdw/edit?usp=share_link)

## Pre-requisites
- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/manuals/)
- [Docker-compose](https://docs.docker.com/compose/)
- [Swagger](https://github.com/swaggo/gin-swagger)

## Getting started
To run the application, make sure you have installed:
- Go
- Docker 
- Docker-compose
- Swagger

If you don't have some of these you can face some errors, so make sure you have all of them installed, following the instructions at [Pre-requisites](#pre-requisites) section.

1. Clone the repository
```shell
git clone https://github.com/murilogilfelpeto/ps-tag-onboarding-go
```

2. Navigate to ps-tag-onboarding-go
```shell
cd ps-tag-onboarding-go
```

3. Run the command:
```shell
make
```

## Test application
Open your browser and access http://localhost:8080/swagger/index.html

## Verify database
Open your browser and access http://localhost:8081
user: admin
pass: pass