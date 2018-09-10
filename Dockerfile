FROM golang

COPY gotask bin

WORKDIR gotask

COPY . .

ENTRYPOINT [ "gotask" ]