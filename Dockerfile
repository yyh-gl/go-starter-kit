FROM golang:1.14.2 AS build
WORKDIR /go/src/api
ENV TZ="Asia/Tokyo"
ENV GO111MODULE=on
COPY . .
RUN go get -u github.com/golangci/golangci-lint/cmd/golangci-lint \
    && go get -u github.com/cosmtrek/air
RUN make build

FROM alpine:latest AS app
COPY --from=build /go/src/api/cmd/api/bin/server /usr/local/bin/server
EXPOSE 8080
ENTRYPOINT ["server"]
