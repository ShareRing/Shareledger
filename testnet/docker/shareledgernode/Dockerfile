FROM ubuntu:18.04

RUN apt-get update && \
    apt-get -y upgrade && \
    apt-get -y install curl wget jq file build-essential libsnappy-dev && \
    wget https://github.com/google/leveldb/archive/v1.20.tar.gz && \
        tar -zxvf v1.20.tar.gz && \
        cd leveldb-1.20/ && \
        make && \
        cp -r out-static/lib* out-shared/lib* /usr/local/lib/ && \
        cd include/ && \
        cp -r leveldb /usr/local/include/ && \
        ldconfig && \
        rm -f v1.20.tar.gz

WORKDIR /shareledger
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

COPY wrapper.sh /usr/bin/wrapper.sh

