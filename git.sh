#!/bin/bash
# @Ralph|2020.09.04


set -x

GOPATH="/go/src/github.com/go"
touch ${GOPATH}/GIT.txt
export Save_Git_Branch_Path="${GOPATH}/GIT.txt"
git branch > ${Save_Git_Branch_Path}
Git_Branch = `awk -F " " '{print $2}' "${Save_Git_Branch_Path}" `
echo "Get git brach finish....."

echo "Push code ....."
git add . 

git commit -m "[go][ralph.Build][1/1][BugId:N/A][无风险][无依赖][自然集成]{$1}"

git push

rm -rf ${GOPATH}/GIT.txt
