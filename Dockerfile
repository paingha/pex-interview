# select image
FROM golang:1.15-alpine
WORKDIR /app
COPY ./go.mod ./go.sum ./

RUN go mod download
ADD . .
WORKDIR /app
RUN go mod vendor
RUN go build -o server
# Command to run the executable
CMD ["./server"]
