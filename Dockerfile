FROM billyteves/alpine-golang-glide:latest

MAINTAINER ndimatulac

COPY ./check-tls-client.go /usr/local/bin/check-tls-client.go

RUN go build -o /usr/local/bin/check-tls-client /usr/local/bin/check-tls-client.go

RUN chmod +x /usr/local/bin/check-tls-client

ENTRYPOINT ["check-tls-client"]
