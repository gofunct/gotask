FROM golang:alpine3.7
RUN apk add --update \
        sqlite       \
        git         \
        build-base
RUN go get -u \
        github.com/Masterminds/glide \
        github.com/gofunct/gotask
WORKDIR /go/src/github.com/gofunct/gotask
RUN glide install
EXPOSE 8080
RUN cat schema.sql | sqlite3 tasks.db 
RUN go install
ENTRYPOINT [ "gotask" ]

