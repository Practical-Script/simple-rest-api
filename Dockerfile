FROM golang:1.22 as build

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-rest-api

FROM golang:1.22-alpine

WORKDIR /app
COPY --from=build /simple-rest-api /app/simple-rest-api

EXPOSE 3000

CMD ["/app/simple-rest-api"]
