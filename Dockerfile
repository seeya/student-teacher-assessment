FROM golang:1.20.2-alpine3.16

WORKDIR /app

RUN apk add --update make

COPY . .
 
RUN go mod download
 
RUN go build -o /godocker

CMD [ "/godocker" ]