FROM golang:1.13

WORKDIR /go/src/github.com/itskass/headhunter
COPY . .

RUN go get -d -v ./...
RUN go build -i -o hhunt .

ENTRYPOINT ["/go/src/github.com/itskass/headhunter/hhunt"]