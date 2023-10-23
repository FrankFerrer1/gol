FROM golang:1.21.1

WORKDIR /app

RUN mkdir -p /app/bin

COPY src/ /app/src

WORKDIR /app/src

RUN go get
RUN go build -o /app/bin/ngp-server /app/src/*.go

ENTRYPOINT ["/app/bin/ngp-server"]
