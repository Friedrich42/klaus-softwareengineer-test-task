# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19-buster AS build

WORKDIR /klaus-softwareengineering-test-task

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /klaus-softwareengineering-test-task

EXPOSE 8080
EXPOSE 8081

USER nonroot:nonroot

COPY database.db .

COPY --from=build /klaus-softwareengineering-test-task/main .

ENTRYPOINT ["/klaus-softwareengineering-test-task/main"]
