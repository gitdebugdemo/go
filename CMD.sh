#!/usr/bin/env bash
cd `dirname $0`
#输出错误和成功信息的函数,带底色
function error(){ echo  -e "\033[41;36m $1 \033[0m";}
function ok(){ echo  -e "\033[42;30m $1 \033[0m";}
if [ -f ./before.sh ]; then
  source ./before.sh
fi
ok "开始运行测试"
#测试服务结果
/usr/local/go/bin/go test -mod=vendor ./test

ok "测试结束"
if [ -f ./after.sh ]; then
  source ./after.sh
fi