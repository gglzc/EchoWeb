FROM golang:1.18.4
WORKDIR /app
ADD . /app
RUN go run main.go
EXPOSE 8080

CMD [ "/build" ]
