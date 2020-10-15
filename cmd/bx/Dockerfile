FROM golang:1.10.4 

RUN apt-get update
RUN apt-get install vim -y
RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get install -y migrate

WORKDIR /go/src/github.com/demsasha4yt/bx-backend-trainee-assignment
