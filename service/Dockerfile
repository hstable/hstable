FROM golang:alpine
ENV GOPROXY=https://goproxy.io
WORKDIR /service
ADD . .
RUN go build -o hstable .
RUN ln ./hstable /usr/bin/hstable
CMD hstable
