FROM golang:1.23-alpine3.21

WORKDIR /go/src

# hot reload
RUN go install github.com/cosmtrek/air@v1.52.0

ENV ENV_PATH=.env

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

CMD ["air", "-c", "./.air.toml"]

