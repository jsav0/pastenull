FROM golang:alpine as builder
WORKDIR /usr/local/go/src/pastenull
COPY main.go .
COPY go.mod .
RUN go build && go install

FROM alpine:latest
EXPOSE 1337
VOLUME /srv
WORKDIR /srv
COPY --from=builder /usr/local/go/bin/pastenull .
ENTRYPOINT ["./pastenull"] 
