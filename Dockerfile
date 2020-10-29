FROM golang:1.14 as builder

RUN mkdir -p /go/src/promadg

WORKDIR /go/src/promadg

COPY . .

RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o promadg .

FROM alpine:latest
WORKDIR /
COPY --from=builder /go/src/promadg/promadg .
