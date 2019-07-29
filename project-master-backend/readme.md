# Standard Go

Standard project microservice for Gophers.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development.

### Installation

First create folder like this `{GOPATH}/src/gitlab.com/standard-go`. 

```
cd {GOPATH}/src/gitlab.com/standard-go/
git clone git@gitlab.com:standard-go/project.git
```

Next,

```
dep ensure -v
cp .env.example .env
go run cmd/web/main.go
```

End with an example of getting some data out of the system or using it for a little demo