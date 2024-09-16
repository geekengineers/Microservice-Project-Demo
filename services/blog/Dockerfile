FROM golang:1.22.4-alpine

WORKDIR /server

COPY . .

RUN make build

CMD [ "./main" ]
