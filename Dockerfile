FROM ubuntu:lunar-20230128

# Install dependencies.

ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

RUN apt-get update
RUN apt-get -y install vim wget
RUN wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz
RUN tar -xzf go1.20.1.linux-amd64.tar.gz


WORKDIR /app
