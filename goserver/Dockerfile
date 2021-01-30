FROM golang:alpine
RUN apk update && apk add git
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promauto
RUN go get github.com/prometheus/client_golang/prometheus/promhttp

RUN mkdir /app

COPY . /app

WORKDIR /app

EXPOSE 2112

CMD ["go","run","hello.go"]



