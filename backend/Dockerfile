FROM golang:1.25.1-alpine3.22 AS build

WORKDIR /app

COPY . .

RUN go build .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/starter-template .

EXPOSE 8081

CMD ["./starter-template"]
