FROM golang


WORKDIR /go/src/app
COPY . .

RUN go build -o app

EXPOSE 8080
CMD ["app"]