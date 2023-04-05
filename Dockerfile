FROM golang:1.20.2-alpine3.16

RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest

#COPY go.mod go.sum ./
WORKDIR /go/src
RUN mkdir apitesting

#COPY *.go /go/src
#COPY ./ /go/src/testgo
COPY ./ /go/src/apitesting
WORKDIR /go/src/apitesting
RUN go get github.com/spf13/viper
RUN go mod tidy
RUN go mod vendor
CMD ["ginkgo","-r","--keep-going"]