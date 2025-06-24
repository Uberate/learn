+++
date = '2025-06-24T22:57:57+08:00'
draft = false
title = 'Bash boom(Bash 炸弹)'
categories = ["技术"]
tags=["bash", "ln"]
+++

## Bash 炸弹

{{< alert icon="fire" cardColor="#e63946" iconColor="#1d3557" textColor="#f1faee" >}}
注意：本篇文章中的指令可能会导致服务器受损。请勿在生产环境中执行。
{{< /alert >}}

请看如下一行指令:

```bash
:(){:|:&};:
```

如果你在 Linux 执行这个指令，并且没有调整过终端的最大线程数的话，大概在几秒之后，Linux 将无法响应任何请求。
此时基本只能通过宿主机管理或物理重启来恢复。

## 原理

Bash 中，对方法的定义如下:

```bash
function () {
    echo "hello"
}
# 通过 function() 进行调用
function()
```

而所谓的 Bash 炸弹则是定义了一个名为 `:` 的函数。
其函数体包括调用  `:` 和管道 `:` 的逻辑。
同时使用 '&' 标记为后台运行。

最后通过 `:()` 引爆。

本质上是不断的创建线程压榨内存，最后导致 Linux 在出发内存回收之前就将内存耗尽导致 Linux 卡死。

## 预防
可以设置用户的最大线程数来防止机器卡死：
```bash
uname -u 2048
```
即可。

### 但是还需要手动清理
经过实际测试，一旦引爆 bash 炸弹，即使配置了 `uname -u 2048` 仍会导致线程爆炸，而且无法停止。

经过一些尝试，我发现 `ps` 指令已经无法使用，`top` 也无法正确反映出触发炸弹的实际来源。

最后通过 `pstree` 停止了所有终端完成了问题清理。不过很显然，那些没有问题的终端任务也被杀死了。

```bash
kill -9 `pstre -p | grep zsh | grep -v sshd | grep -v node | awk -Fzsh '{print $2}' | sed 's/[()]//g'`
```

反复执行几次后，就会发现系统负载明显降低。

## 参考
[https://zhuanlan.zhihu.com/p/685820978](https://zhuanlan.zhihu.com/p/685820978)
[https://www.jianshu.com/p/9e508888e2d9](https://www.jianshu.com/p/9e508888e2d9)