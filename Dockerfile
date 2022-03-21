FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get github.com/githubnemo/CompileDaemon     

EXPOSE 8080

CMD ["go run main.go"]