FROM golang:alpine

COPY ./ ./
WORKDIR /cmd
RUN go build -o main .
CMD ["./main"]


