FROM golang:1.17

RUN mkdir /app

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

ADD ./ /app


RUN cd ./app && go build -o api

CMD cd ./app && ./api