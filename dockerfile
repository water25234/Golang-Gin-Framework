FROM golang:latest
WORKDIR /src/github.com/water25234/Golang-Gin-Framework
COPY . /src/github.com/water25234/Golang-Gin-Framework
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./Golang-Gin-Framework"]