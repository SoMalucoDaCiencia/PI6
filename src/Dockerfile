FROM golang AS builder
# RUN apk add git

ENV DOCKER=true

WORKDIR /go/src
COPY . .

ENV DB_SCHEMA=famarcias
ENV DB_USER=famarcias
ENV DB_PASSWORD=famarcias
ENV DB_ADDRESS=www.thyagoquintas.com.bt
ENV DB_PORT=1433


RUN go mod download
RUN go build -o /go-app-pi6 ./core/

RUN chmod +x /go-app-pi6
CMD [ "/go-app-pi6" ]
