FROM golang:1.17.8-alpine AS build
RUN apk add --no-cache git
RUN git clone https://github.com/hejiangda/CloudNativeGeekTime.git
WORKDIR /go/CloudNativeGeekTime/lesson2/homework1
ENV GO111MODULE=auto
RUN go build -o /bin/httpserverHomework

FROM alpine
COPY --from=build /bin/httpserverHomework /bin/httpserverHomework
ENTRYPOINT ["/bin/httpserverHomework"]