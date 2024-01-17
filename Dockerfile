FROM golang:alpine

LABEL ese.image.author="saumyabhatt10642"

ENV GIN_MODE=release

WORKDIR /src

COPY ./src .

RUN go get .
RUN go build -o server .

EXPOSE $SERVER_PORT

CMD [ "./server" ]
