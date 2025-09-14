+++
date = '2025-09-14T01:09:13+08:00'
title = '读《Your biggest customer might be your biggest bottleneck》'
categories = ["阅读"]
tags = ["短篇阅读"]
+++

> 原文 [Your biggest customer might be your biggest bottleneck](https://densumesh.dev/blog/fair-queue/) -- [densumesh](https://github.com/densumesh)

吵闹邻居，
是指在多租户隔离模式下，
共享资源被独立大租户过度消耗而导致的系统整体异常的问题。

作者之前使用的 FIFO 的队列模型由于某个租户的巨量输入而出现问题。
然后作者提出了一个新的方案。
轮询所有通道队列，依次处理。

---

我觉得这个模式通过 Golang 比较好 Mock，所以开个仓库鼓捣一下试试看。

{{< github repo="Uberate/fair_go_chan_mock" >}}