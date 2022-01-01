FROM golang:1.17-alpine

RUN mkdir /app
WORKDIR /app

COPY . .

RUN apk add make && make setup && make build

CMD ["make", "start-prod"]