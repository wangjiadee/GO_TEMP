#!/bin/bash
# @Ralph|2020.09.04
# 使用脚本携带2个参数 脚本所在的绝对路径 commit内容


set -x


DIYGOPATH= $1
touch ${DIYGOPATH}/GIT.txt
export Save_Git_Branch_Path="${DIYGOPATH}/GIT.txt"
git branch > ${Save_Git_Branch_Path}
Git_Branch = `awk -F " " '{print $2}' "${Save_Git_Branch_Path}" `
echo "Get git brach finish....."

echo "Push code ....."
git add --all . 

git commit -m "[go][ralph.Build][1/1][BugId:N/A][无风险][无依赖][自然集成]{$2}"

git push

rm -rf ${DIYGOPATH}/GIT.txt
