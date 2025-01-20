FROM golang:1.23.4-alpine as build

WORKDIR /go-app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN apk add --no-cache make
RUN make build

FROM alpine:latest

RUN apk add --no-cache postgresql-client

COPY --from=build /go-app/bin/app/fissionNode-backend /bin/app/fissionNode-backend

COPY --from=build /go-app/.env /bin/app/.env

EXPOSE 8080

CMD ["/bin/app/fissionNode-backend"]
