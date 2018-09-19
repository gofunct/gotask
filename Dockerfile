FROM golang

RUN apt update && sudo apt upgrade

RUN apt-get install sqlite3

COPY gotask bin

WORKDIR gotask

COPY . .

ENTRYPOINT [ "gotask" ]