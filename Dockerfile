FROM golang

RUN apt update -y && apt upgrade -y

RUN apt-get install sqlite3

COPY gotask bin

WORKDIR gotask

COPY . .

ENTRYPOINT [ "gotask" ]