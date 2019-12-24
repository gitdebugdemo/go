#!/usr/bin/env bash
cd `dirname $0`
ok "开始运行测试"
#测试服务结果
/usr/local/go/bin/go test ./test
