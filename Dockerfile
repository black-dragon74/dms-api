# Stage 1: Builds the source
FROM golang:1.16-alpine as buildstage

WORKDIR /usr/src/app

# Copy Source
COPY . .

# Build
RUN go build -o api-server main.go

# Stage 2: Just executes the source
FROM alpine as prodstage

WORKDIR /usr/src/app

# Copy the built binary
COPY --from=buildstage /usr/src/app/api-server .
COPY --from=buildstage /usr/src/app/config.toml .

# Change the dev environment to production
RUN sed -i 's/dev/prod/g' ./config.toml

# Change interface from loopback to broadcast, thx @sid-sun
RUN sed -i 's/localhost/0.0.0.0/g' ./config.toml

# Off we go!
CMD ["./api-server"]
