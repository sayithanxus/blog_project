FROM golang:alpine3.12
WORKDIR /go/src/goweb
COPY ../../OneDrive/Masaüstü/goblog-demo-main .
CMD ["/go/src/goweb/goblog"]