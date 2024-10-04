FROM golang:1.22.5-alpine

RUN apk update \
    && apk add --no-cache gcc \
    && apk add --no-cache build-base

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o ./pastebin

EXPOSE 8000

CMD ["./pastebin"]