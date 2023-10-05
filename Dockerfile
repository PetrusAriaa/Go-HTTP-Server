FROM golang:1.21.1

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy

RUN go build -o ./out/dist .

ENV MONGODB_URI=$MONGODB_URI
RUN echo $MONGODB_URI


CMD ./out/dist