+++
date = '2025-09-14T17:02:26+08:00'
title = '公平竞争 Chan 模型'
categories = ['技术']
tags = ['architecture']
+++

源自这个文章：
{{< article link="/learn/posts/article/short/11_your_biggest_customer_might_be_your_biggest_bottleneck/" showSummary=true compactSummary=true >}}


仓库：
{{< github repo="Uberate/fair_go_chan_mock" >}}

公平竞争模型在 Golang 中似乎非常好模拟.

我们先来看一下 golang 中如何实现一个基础的 FIFO 模型.

### 消息

整个模拟我希望可以简单一点，
所以我们简化一下消息内容：

```golang
type Message struct{
    Tenant string // 租户信息
    Message int   // 消息内容
}
```

### 队列

> https://github.com/Uberate/fair_go_chan_mock/blob/main/type.go

{{< codeimporter url="https://raw.githubusercontent.com/Uberate/fair_go_chan_mock/refs/heads/main/type.go" type="go" >}}

### FIFO 

这个对于 Golang 来说非常简单，直接用一个 `chan` 即可

> https://github.com/Uberate/fair_go_chan_mock/blob/main/fifo.go

{{< codeimporter url="https://raw.githubusercontent.com/Uberate/fair_go_chan_mock/refs/heads/main/fifo.go" type="go" >}}

### Fair 模型

这个对于 golang 来说，也不是特别困难，这里直接给每个租户都添加了一个 Chan 。

> https://github.com/Uberate/fair_go_chan_mock/blob/main/fair.go

{{< codeimporter url="https://raw.githubusercontent.com/Uberate/fair_go_chan_mock/refs/heads/main/fair.go" type="go" >}}

## Mock
然后我们就来模拟一下吧~

我们构造这个输入:
```text
t1: +++++.++++++++++++....+++++++++++++
t2: +...+.............+...+............
t3: ....++............+++++............
```

t1、t2、t3 表示三个不同的租户。上面的代码段中，`+` 和 `.` 分别表示输入的状态。

然后我们运行后，得到了如下两个输出结果

FIFO:
```text
t1: ++.+++...++++++++++++.....+..++++++++++++
t2: ..+...+.................+..+.............
t3: .......++............+++.+..+............
```

Fair:
```text
t1: +.+..+..+..+.+.+.+.++++++++++++++++++++++
t2: .+.+..+..+...............................
t3: ....+..+..+.+.+.+.+......................
```

可以看到 FIFO 的 t2、t3 终止在相对较后的位置，但是在 Fair 的模式下，较早的结束了。

事实证明，通过 Fair 的逻辑可以解决邻居吵闹的模型。