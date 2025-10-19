+++
date = '2025-10-19T18:12:06+08:00'
title = '读《URL Design》'
categories = ['阅读']
tags = ['短篇阅读']
+++

> 原文 [URL Design](https://warpspire.com/posts/url-design) -- Kneath


一开始以为是面向 OpenAPI 的设计，
后来发现是面向网站的设计。
这篇文章写在 2010 年，
距今已有 15 年的时间了，
15 年中，
互联网飞速发展，
URL 已经不再是网站的专属了。
现在的大型服务，
网站等，
都在通过 URL 来快速交换信息。

与此同时，
HTTP 协议快速演进，
在此之上还诞生了众多的 RPC 远程调用。
所有基于 HTTP 的服务，
实际都是在以来 URL。

{{< alert "bell" >}}
所以，
文中有一点我是非常认可的，
**URL 是面向人的**。
这在现在仍然不过时。
{{< /alert >}}

不过目前的可观测技术来说，
大多数的请求都要使用 Trace 等技术，
有时候的确会有人通过 QueryParam 来传递请求信息。
所以在这种情况，
URL 最起码需要保证 Path 部分是人类可读的。

而目前的 URL Path 设计，
也诞生了非常多的优秀规范。
比如 Restful 等风格。



