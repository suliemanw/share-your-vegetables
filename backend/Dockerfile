FROM golang:1.25.1-alpine3.22 AS build

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o share-your-vegetables

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/share-your-vegetables .

COPY --from=builder /app/uploads ./uploads

EXPOSE 8081

CMD ["./share-your-vegetables"]
