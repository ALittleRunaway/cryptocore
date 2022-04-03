FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build -o cryptocore

EXPOSE 6001

CMD [ "./cryptocore" ]