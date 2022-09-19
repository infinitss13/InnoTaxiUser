FROM golang:alpine

RUN apk update
RUN apk add postgresql-client


WORKDIR /app

COPY ./ /app


RUN go mod download

#RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz
#RUN tar -xzf migrate.linux-amd64.tar.gz
#RUN mkdir -p ~/bin
#RUN mv migrate.linux-amd64 ~/bin/migrate
#RUN migrate -path ./schema -database 'postgres://postgres:qwerty@postgres:5432/postgres?sslmode=disable' up
#RUN rm migrate.linux-amd64.tar.gz

#RUN  go get -u github.com/golang-migrate/migrate/v4
#RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
#RUN migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up


ENTRYPOINT go run cmd/main.go

EXPOSE 8000