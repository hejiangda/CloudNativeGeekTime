FROM golang:1.18-alpine AS build
COPY httpserver.go /go/CloudNativeGeekTime/lesson8/
WORKDIR /go/CloudNativeGeekTime/lesson8/
ENV GO111MODULE=auto
RUN go build -o /bin/httpserverHomework

FROM alpine
COPY --from=build /bin/httpserverHomework /bin/httpserverHomework
EXPOSE 8080
ENTRYPOINT ["/bin/httpserverHomework"]