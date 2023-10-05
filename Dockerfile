FROM golang:1.21.1

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy

RUN go build -o ./out/dist .

RUN echo $MONGODB_URI

ENV MONGODB_URI=${MONGODB_URI}

CMD ./out/dist