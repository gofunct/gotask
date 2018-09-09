FROM golang
COPY gotask $GOPATH/bin
WORKDIR /home/gotask
COPY . .
ENTRYPOINT ["gotask"]
