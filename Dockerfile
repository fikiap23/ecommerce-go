FROM golang:1.22.3-alpine

ENV GOFLAGS="-buildvcs=false"

RUN apk add --no-cache git curl
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
CMD ["air"]
