FROM golang:1.21.1

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy

ARG MONGODB_URI
ENV MONGODB_URI ${MONGODB_URI}

RUN go build -o ./out/dist .

CMD ./out/dist