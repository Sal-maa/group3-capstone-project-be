FROM golang:1.17

RUN mkdir /app

WORKDIR /app

COPY .env /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download

ADD ./ /app

RUN cd ./app && go build -o api

CMD cd ./app && ./api