FROM alpine:3.7

RUN apk update && \
    apk upgrade && \
    apk --no-cache add curl jq file

VOLUME  /shareledger
WORKDIR /shareledger
EXPOSE 1317
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["rest-server"]
STOPSIGNAL SIGTERM

COPY wrapper.sh /usr/bin/wrapper.sh

