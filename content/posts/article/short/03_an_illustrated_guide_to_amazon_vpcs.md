+++
date = '2025-06-15T18:55:14+08:00'
title = '读 《An illustrated guide to Amazon VPCs》'
categories = ["阅读"]
tags = ["短篇阅读", "网络", "VPC"]
+++

> 原文 [An illustrated guide to Amazon VPCs](An illustrated guide to Amazon VPCs) -- Aditya Bhargava 

作者阐述了 AWS 中 `VPC` 的诞生原因：
- 在 AWS 中，每个用户都有自己的私有网络，称作 `VPC`。
- 如果没有专用网络，IP 就会出现地址冲突的问题。
- 如果没有专用网络，每个人都在同一个网络下，这不安全。
- VPC 是通过映射服务（Mapping service）实现的。
  
`VPC` 的全称: `Virtual Private Cloud`。

一般来说，VPC 应该会架设在 `ISO 网络七层` 中的三层（网络层）。

不了解网络七层的朋友可以看这篇文章，我大概标记了一下网络七层的作用:
[网络基础](/learn/posts/tech/network/base/)