FROM golang

WORKDIR /InnoTaxiUser
COPY ./ ./

RUN go build -o /docker-gs-ping

CMD [ "/docker-gs-ping" ]



