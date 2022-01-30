# =============================================================================
#
#
# Project:
#     https://github.com/jotka/ac_rest_go.git
#     Home Asisstant IP climate middleware for Samsung AC devices.
#
#
# =============================================================================
#
##
## Build
##
FROM golang:1.16-buster AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY in ./
COPY out ./
RUN go build -o /ac-rest

##
## Deploy
##
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /ac-rest /ac-rest
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/ac-rest"]


