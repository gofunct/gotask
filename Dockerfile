FROM golang

RUN sudo apt update && sudo apt upgrade

RUN sudo apt-get install sqlite3

COPY gotask bin

WORKDIR gotask

COPY . .

ENTRYPOINT [ "gotask" ]