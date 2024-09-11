FROM golang AS builder
# RUN apk add git

ENV DOCKER=true

WORKDIR /go/src
COPY . .

RUN go mod download
RUN go build -o /go-app-pi6 ./core/

RUN chmod +x /go-app-pi6
CMD [ "/go-app-pi6" ]
