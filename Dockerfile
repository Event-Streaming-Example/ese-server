FROM golang:alpine

LABEL ese.image.author="saumyabhatt10642"

ENV GIN_MODE=release

RUN mkdir /app
COPY src/ /app
WORKDIR /app

RUN go get .
RUN go build -o server .

CMD [ "/app/server" ]
