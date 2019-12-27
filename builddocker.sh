#!/usr/bin/env bash

docker build -t  registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e05c2b1c04c51577435825 ./;

docker push registry.cn-shanghai.aliyuncs.com/copygithub/github:gitdebugdemo-5e05c2b1c04c51577435825;

