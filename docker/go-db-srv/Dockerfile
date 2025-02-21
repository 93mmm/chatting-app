FROM golang:1.23.5-alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache make

CMD [ "make", "pgserver" ]
