FROM golang:1.21.1-alpine3.18 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

# Final stage
FROM scratch

COPY --from=build /app/app /

CMD ["/app"]
