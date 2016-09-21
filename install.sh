#!/usr/bin/env bash
#之所以加上这个install，是不用配置GOPATH（避免新增一个GO项目就要往GOPATH中增加一个路径）
if [ ! -f install ]; then
echo 'install must be run within its container folder' 1>&2
exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src

go install test

export GOPATH="$OLDGOPATH"

echo 'finished'