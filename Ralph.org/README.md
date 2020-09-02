# go


go commit 提交格式：
[go][ralph.Build][1/1][BugId:N/A][无风险][无依赖][自然集成]{有关内存的理解}

风险及影响[快/稳/省/功能]：{无}
测试建议：{无需测试}
跨组依赖(提交链接)：{无}
移植适用范围：{无}



# src/github.com/go/Ralph.org/test/run/main.go:6:2: cannot find module providing package github.com/go/Ralph.org/test/rc: working directory is not part of a module

解决方案：export GO111MODULE=auto

# go包环境问题

设置代理： GOPROXY="https://goproxy.io,direct"