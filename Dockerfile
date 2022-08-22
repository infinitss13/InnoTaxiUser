FROM golang

WORKDIR /InnoTaxiUser
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o inno



