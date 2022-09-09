FROM golang:1.18.2-stretch AS build

WORKDIR /app

RUN apt-get update && apt-get --assume-yes upgrade make

COPY /go.mod ./
COPY /go.sum ./
RUN go mod tidy
RUN go mod download

COPY  . .
RUN make build

FROM alpine:3.14
WORKDIR /app

RUN apk add --no-cache libc6-compat

COPY --from=build /app/build/* ./
RUN chmod +x ./shareledger
RUN mv ./shareledger /bin
#
EXPOSE 26656 26657
CMD ["shareledger", "start"]