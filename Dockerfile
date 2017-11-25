FROM jaehue/golang-onbuild
MAINTAINER jang.jaehue@eland.co.kr

# install go packages
RUN go get github.com/relax-space/lemon-wxpay-sdk && \
    go get github.com/relax-space/go-kit/...


# add application
ADD . /go/src/lemon-epay-api
WORKDIR /go/src/lemon-epay-api
RUN tar xf tmp/wxcert.tar.gz -C /go/src/github.com/relax-space/lemon-wxpay-sdk
RUN go install

EXPOSE 5000

CMD ["lemon-epay-api"]