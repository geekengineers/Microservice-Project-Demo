FROM golang:1.22.4

WORKDIR /server

COPY . .

RUN go mod tidy

ENV GO_ENV=production

RUN go build -o blog_server ./services/blog/cmd/grpc/main.go

CMD [ "./blog_server" ]
