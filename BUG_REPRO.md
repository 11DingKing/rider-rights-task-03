# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

服务站做超期提醒时，承诺期限正好走到截止时刻的事项没有出现在列表里，晚一分钟的事项却能被找到。先不要修改代码，沿着这次筛选找到涉及的 Go 文件和符号，说明边界为何漏掉事项，并用定向和全量测试结果把诊断落下来。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-03
- 仓库地址：https://github.com/11DingKing/rider-rights-task-03.git
- parent SHA：e49dc9c163248c6da821f181304689b15cb0e8db

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-03.git bug-repro
cd bug-repro
git checkout --detach e49dc9c163248c6da821f181304689b15cb0e8db
go test ./internal/domain -run "^TestDeadlineMomentIsOverdue$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestDeadlineMomentIsOverdue$" -count=1
--- FAIL: TestDeadlineMomentIsOverdue (0.00s)
    task03_test.go:12: case at deadline was not treated as overdue
FAIL
FAIL	riderguard/internal/domain	0.039s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestDeadlineMomentIsOverdue$" -count=1
--- FAIL: TestDeadlineMomentIsOverdue (0.00s)
    task03_test.go:12: case at deadline was not treated as overdue
FAIL
FAIL	riderguard/internal/domain	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

先不要修改目标仓库代码。通过源码定位和实际定向验证说明具体文件、符号以及等于截止时间时返回 false 如何导致超期提醒缺失；诊断结论必须与源码和测试证据一致，定向验证、相关包测试和仓库全量回归命令的结果必须与结论一致；诊断结束时目标仓库代码、测试和配置零改动，不得删除、跳过或削弱测试。
