FROM golang:1.22.4-alpine

WORKDIR /server

COPY . .

RUN apt update

RUN apt install make

RUN make build

CMD [ "./main" ]
