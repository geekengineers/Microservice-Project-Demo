FROM golang:1.22.4

WORKDIR /server

COPY . .

RUN go mod tidy

ENV GO_ENV=production

RUN go build -o auth_server ./services/auth/cmd/grpc/main.go

CMD [ "./auth_server" ]
