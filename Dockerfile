FROM golang:1.16-alpine as form3_cli
RUN apk add build-base
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8080

#FROM form3_client as test
RUN go test -v -json -c -o /form3_cli_test
CMD ["/form3_cli_test", "-test.v"]

#FROM form3_client as build
RUN go build -o /form3_cli .
#CMD ["/form3_cli"]





