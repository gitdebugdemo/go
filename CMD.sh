#!/bin/bash
#初始化mysql-root账户
export MYSQL_ROOT_PASSWORD="123456"
export MYSQL_DATABASE="demo"
export MYSQL_USER="demo"
export MYSQL_PASSWORD="demo"

#输出错误和成功信息的函数,带底色
function error(){ echo  -e "\033[41;36m $1 \033[0m";}
function ok(){ echo  -e "\033[42;30m $1 \033[0m";}

#观察mysql的端口是否已经起来了
mysql_up=$(netstat -tnlpu | grep 3306 | wc -l)
ok "mysql_up: $mysql_up"
if [ $mysql_up -gt 0 ]; then
  ok "mysql 服务启动成功!"
else
    #启动mysql服务
    docker-entrypoint.sh mysqld &

    #等待20秒.等mysql服务器启动之后,才执行业务
    for times in $(seq 1 300)
    do
      #观察mysql的端口是否已经起来了
      mysql_up=$(netstat -tnlpu | grep 3306 | wc -l)
      ok "mysql_up: $mysql_up"
      if [ $mysql_up -gt 0 ]; then
        ok "mysql 服务启动成功!"
        break ;
      fi
      ok "$times"
      sleep 1
    done
    if [ $mysql_up -eq 0 ]; then
        error "mysql 服务启动失败[n]"
        exit 1
    fi
fi
#启动go语言本身服务
cd `dirname $0`

#环境准备,比如数据库,接口启动等
if [ -f ./before.sh ]; then
  source ./before.sh
fi

if [ -f /root/main.go ]; then
  ok "启动go主程序"
  #启动go服务
  /usr/local/go/bin/go run -mod=vendor /root/main.go &
fi

ok "开始运行测试"
#测试服务结果
/usr/local/go/bin/go test -mod=vendor ./test -v
ok "测试结束"
if [ -f ./after.sh ]; then
  source ./after.sh
fi