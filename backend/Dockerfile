FROM golang:1.24.3-alpine3.21

RUN mkdir /app
ADD . /app 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o main .
CMD ["/app/main"]
