FROM golang:1.13.5-stretch

WORKDIR  /root/

#git源码拷贝到运行目录下
COPY /  /root/

CMD ["/usr/local/go/bin/go","run","/root/src/main.go"]