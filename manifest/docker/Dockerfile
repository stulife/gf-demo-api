
FROM golang:alpine
WORKDIR /data
ADD . ./

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o app .


FROM alpine:latest

###############################################################################
#                                INSTALLATION
###############################################################################

WORKDIR /home/data
COPY --from=0 /data/app  /home/data

COPY --from=0 /data/manifest  /home/data/manifest
COPY --from=0 /data/resource  /home/data/resource


EXPOSE 8000

ENTRYPOINT  ["./app"]
