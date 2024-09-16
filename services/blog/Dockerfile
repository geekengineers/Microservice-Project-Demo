FROM golang:1.22.4-alpine

WORKDIR /server

COPY . .

RUN apk update

RUN apk add make

RUN go mod tidy

RUN make build

CMD [ "./main" ]
