+++
date = '2025-06-02T19:10:41+08:00'
title = 'GO Channel 的空值特性'
tags = ["ln", "go"]
categories = ["技术"]
author = "Uberate"
topic = "深入 Go"
+++

# Golang Channel nil 

Channel 在空值的时候会有特殊的行为。

与大多数的数据类型不同，Channel 在 `nil` 的情况下进行读写会发生线程阻塞，而不是空指针异常。

## ReadNil

{{< ci url="codes/go/channel/readnil/read_nil_channel.go" >}}

## WriteNil

{{< ci url="codes/go/channel/writenil/write_nil_channel.go" >}}

执行上述两个代码我们发现了一个现象，`nil` 状态下的 Channel 是无法读写的，会一直阻塞。

## 避免

那有什么方式可以避免吗？实际上这个问题是一个 `Go` 陷阱。需要我们在所有 `Read` 或 `Write` 的时候，使用 `select` 关键字。

例如如下代码
{{< ci url="codes/go/channel/resumechannelnil/resume_channel_nil.go" >}}