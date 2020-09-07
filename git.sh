#!/bin/bash
# @Ralph|2020.09.04


set -x
GITFILE = "GIT.txt"
GOPATH="/go/src/github.com/go"
mkdir ${GOPATH}/${GITFILE}
export Save_Git_Branch_Path="${GOPATH}/${GITFILE}"
Git_Branch = `awk -F " " '{print $2}' "${Save_Git_Branch_Path}" `

git branch > ${Save_Git_Branch_Path}

git add . 

git commit -m "[go][ralph.Build][1/1][BugId:N/A][无风险][无依赖][自然集成]{$1}

风险及影响[快/稳/省/功能]：{无}
测试建议：{无需测试}
跨组依赖(提交链接)：{无}
移植适用范围：{无}"

git push

rm -rf ${GOPATH}/GIT.txt
