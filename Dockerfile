FROM golang:alpine3.7
RUN apk add --update \
        sqlite       \
        git         \
        build-base
RUN go get -u \
        github.com/Masterminds/glide \
        github.com/ops2go/gotask
WORKDIR /go/src/github.com/ops2go/gotask
EXPOSE 8081
RUN cat schema.sql | sqlite3 tasks.db
RUN go build
ENTRYPOINT [ "gotask" ]

