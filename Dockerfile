FROM golang:1.20.4-buster AS build

WORKDIR /app

# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY *.go ./
COPY . ./

RUN ls

RUN go build -o app .
RUN chmod +x app

CMD ["./app"]