FROM golang:1.8

RUN apt-get update -y && apt-get install -y zip unzip git autoconf automake libtool

RUN cd /opt && \
    git clone https://github.com/google/protobuf -b v3.4.1 --depth 1 && \
    cd /opt/protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install

RUN go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/duhruh/blog

COPY . .

RUN glide install
