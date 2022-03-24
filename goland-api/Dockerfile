FROM golang:1.17.8

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /run.app

EXPOSE 8080

ENTRYPOINT [ "/run.app" ]
