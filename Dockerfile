FROM golang:1.21.1

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy

ENV MONGODB_URI=${MONGODB_URI}

RUN MONGODB_URI=${MONGODB_URI} go build -o ./out/dist .

CMD ./out/dist