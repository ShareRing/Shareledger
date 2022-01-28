FROM golang:1.17-alpine AS build

WORKDIR /app

RUN apk add --update make

COPY /go.mod ./
COPY /go.sum ./
RUN go mod tidy
RUN go mod download

COPY  . .
RUN make build

FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/build/* ./
RUN chmod +x ./shareledger
RUN mv ./shareledger /bin
#
EXPOSE 26656 26657
CMD ["shareledger", "start"]