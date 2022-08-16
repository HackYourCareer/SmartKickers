FROM golang:alpine as builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN apk add gcc musl-dev
RUN go test ./...
RUN go build -o main .

FROM alpine
COPY --from=builder /build/main /app/
EXPOSE 3000
WORKDIR /app
CMD ./main
