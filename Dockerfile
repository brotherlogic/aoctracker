# syntax=docker/dockerfile:1

FROM golang:1.18 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/aoctracker

COPY go.mod ./
COPY go.sum ./

RUN mkdir proto
COPY proto/*.go ./proto/

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /aoctracker

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /aoctracker /aoctracker

EXPOSE 8080
EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/aoctracker"]