FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod tidy


ENTRYPOINT [ "go", "run", "./cmd/main.go" ]