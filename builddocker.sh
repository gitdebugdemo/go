#!/usr/bin/env bash

docker build -t  registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e01989ee37771577162910 ./;

docker push registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e01989ee37771577162910;

