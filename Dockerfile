FROM golang:alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o karmchari .

CMD ["/app/karmchari", "-port", "51463", "-shost", "redis:6379"]

