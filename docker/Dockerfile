# syntax=docker/dockerfile:1

FROM golang:1.25.1 AS build 
WORKDIR /app

# cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source and build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./app

# runtime stage
FROM alpine:3.20
RUN adduser -D -H app && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /app/server /usr/local/bin/server
ENV PORT=8080
EXPOSE 8080
USER app
ENTRYPOINT ["server"]