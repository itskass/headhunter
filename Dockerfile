FROM golang:1.13

WORKDIR /go/src/github.com/itskass/headhunter
COPY . .

RUN go get -d -v ./...
RUN go build -i -o headhunter .

ENTRYPOINT ["/go/src/github.com/itskass/headhunter/headhunter"]