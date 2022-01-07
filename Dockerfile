# select image
FROM golang:1.17-alpine
WORKDIR /app
COPY ./go.mod ./go.sum ./

RUN go mod download
ADD . .
WORKDIR /app
RUN go mod vendor
RUN go build -o main
# Command to run the executable
CMD ["./main"]
