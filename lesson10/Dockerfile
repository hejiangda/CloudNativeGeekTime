FROM golang:1.18-alpine AS build
COPY . /go/CloudNativeGeekTime/lesson10/
WORKDIR /go/CloudNativeGeekTime/lesson10/
ENV GO111MODULE=auto
RUN go mod init httpserver && go mod tidy && go build -o /bin/httpserverHomework

FROM alpine
COPY --from=build /bin/httpserverHomework /bin/httpserverHomework
EXPOSE 8080
ENTRYPOINT ["/bin/httpserverHomework"]