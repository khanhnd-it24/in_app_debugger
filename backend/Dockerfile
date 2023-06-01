# Builder
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /main .

FROM alpine

COPY --from=builder /main /main
COPY ./configs ./configs

EXPOSE 2345

CMD ["/main"]