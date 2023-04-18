FROM golang:1.18.4-alpine3.15 AS build
WORKDIR /app

RUN set -eux; apk add --no-cache ca-certificates build-base linux-headers;
RUN apk add git

COPY /go.mod ./
COPY /go.sum ./
RUN go mod tidy
RUN go mod download
COPY  . .

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.a

RUN make build LEDGER_ENABLED=false BUILD_TAGS=muslc

FROM alpine:3.14
WORKDIR /app

RUN apk add --no-cache libc6-compat

COPY --from=build /app/build/* ./
COPY --from=build /lib/libwasmvm_muslc.a /lib/libwasmvm_muslc.a

RUN chmod +x ./shareledger
RUN mv ./shareledger /bin
#
EXPOSE 26656 26657 26658 1317 9090 9091
CMD ["shareledger", "start"]
