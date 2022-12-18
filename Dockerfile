FROM golang:1.19-alpine

WORKDIR /app

ENTRYPOINT ["tail", "-f", "/dev/null"]