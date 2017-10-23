FROM golang:1.8

ENV APP_PKG_PATH=/go/src/github.com/duhruh/blog

RUN curl https://glide.sh/get | sh

RUN apt-get update -y && apt-get install -y zip unzip git autoconf automake libtool

RUN cd /opt && \
    git clone https://github.com/google/protobuf -b v3.4.1 --depth 1 && \
    cd /opt/protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install

#RUN go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

RUN mkdir -p $APP_PKG_PATH

WORKDIR $APP_PKG_PATH

VOLUME $APP_PKG_PATH

ADD glide.yaml $APP_PKG_PATH
ADD glide.lock $APP_PKG_PATH

RUN glide install

ADD . $APP_PKG_PATH


