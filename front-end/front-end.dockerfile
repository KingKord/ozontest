# base go image
FROM golang:1.21.4-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

# CGO=0 means that we do not use C libraries
RUN CGO_ENABLED=0 go build -o frontEndApp ./cmd/api

# give frontEndApp an executable flag
RUN chmod +x /app/frontEndApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

#COPY --from=builder /app/migrations /app
COPY --from=builder /app/frontEndApp /app

CMD [ "/app/frontEndApp" ]