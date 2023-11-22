FROM golang:1.19.3

WORKDIR /backEnd

COPY . .

ENV GO111MODULE=on

RUN go mod download

RUN chmod +x /backEnd

WORKDIR /backEnd

RUN go build -o main .

EXPOSE 8090

CMD [ "/main" ]

