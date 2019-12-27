FROM xltxlm/mysql:8.0-go.1.13.5

WORKDIR  /root/

#git源码拷贝到运行目录下
COPY /  /root/

CMD ["/root/CMD.sh"]