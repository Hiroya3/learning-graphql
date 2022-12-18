FROM golang:1.19-alpine

RUN go install github.com/cosmtrek/air@v1.40.4

COPY . /app
WORKDIR /app

CMD ["air", "-c", ".air.toml"]