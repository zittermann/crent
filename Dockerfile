FROM golang:1.21.1-alpine3.18 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

# Final stage
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]