FROM golang:1.21.5-alpine
WORKDIR /go_api
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go-api
EXPOSE 8080
CMD ["/docker-go-api"]
