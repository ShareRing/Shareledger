FROM golang:1.18.4-alpine3.15 AS build
WORKDIR /code

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
WORKDIR /code

RUN apk add --no-cache libc6-compat

COPY --from=build /code/build/* ./
COPY --from=build /lib/libwasmvm_muslc.a /lib/libwasmvm_muslc.a

RUN chmod +x ./shareledger
COPY ./shareledger /usr/bin/

RUN echo "Ensuring binary is statically linked ..." \
  && (file ./build/shareledger | grep "statically linked")