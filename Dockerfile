FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

RUN go get github.com/githubnemo/CompileDaemon     

EXPOSE 8000

# CMD ["go run main.go"]
CMD ["go", "run", "main.go"]
# ENTRYPOINT [ "go","run" ]
# CMD ["main.go"]