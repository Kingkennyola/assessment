FROM golang:1.22-bookworm as base

WORKDIR $GOPATH/src/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /phaidra-assessment main.go

FROM gcr.io/distroless/static-debian12

COPY --from=base /phaidra-assessment .

EXPOSE 8080 9095

CMD ["./phaidra-assessment"]
