FROM golang:1.18-alpine
RUN addgroup -S go && adduser -S go -G go
USER go:go
WORKDIR /home/go
CMD mkdir data
COPY ./maps .
COPY go.mod .
RUN go build -o ./out/maps
ENTRYPOINT ["./out/maps"]
