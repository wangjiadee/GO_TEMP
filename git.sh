#!/bin/bash

set -x

GOPATH="/go/src/github.com/go"
mkdir ${GOPATH}/GIT.txt
export Save_Git_Branch_Path="${GOPATH}/GIT.txt"
Git_Branch = `awk -F " " '{print $2}' "${Save_Git_Branch_Path}" `

git branch > $Save_Git_Branch_Path

git add . 

git commit -m "[go][ralph.Build][1/1][BugId:N/A][无风险][无依赖][自然集成]{strcut 结构体相关}

风险及影响[快/稳/省/功能]：{无}
测试建议：{无需测试}
跨组依赖(提交链接)：{无}
移植适用范围：{无}"

git push

rm -rf ${GOPATH}/GIT.txt

