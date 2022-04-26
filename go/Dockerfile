# ビルドイメージの指定
FROM golang:alpine3.15 as builder
WORKDIR /go/src
COPY ./ ./
RUN go mod tidy
RUN go build -o main

# マルチステージビルド
FROM alpine
# builderからapp
COPY --from=builder /go/src /app
EXPOSE 8080
CMD ["./app/main"]
