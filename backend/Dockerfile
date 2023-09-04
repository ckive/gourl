FROM golang:1.21.0

WORKDIR /api

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /urlshortener_api

EXPOSE 9808

# Run
CMD ["/urlshortener_api"]