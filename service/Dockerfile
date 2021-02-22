FROM golang:alpine
ENV GOPROXY=https://goproxy.io
WORKDIR /service
ADD . .
ENV GIN_MODE=release
RUN go build -o hstable .
RUN ln ./hstable /usr/bin/hstable
RUN echo "10.248.17.16 jw.hitsz.edu.cn" >> /etc/hosts
RUN echo "219.223.236.40 sso.hitsz.edu.cn" >> /etc/hosts
CMD hstable
