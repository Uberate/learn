+++
date = '2025-06-10T23:17:14+08:00'
draft = true
title = '基础概念'
tags = ['AI', 'MCP']
categories = ['技术']
+++

## 什么是 MCP ？
### 定义
MCP (**M**odel **C**ontext **P**rotocol) 是一个模型上下文协议。
由 `Anthropic` 自 `2024/11` 发布后，在 [mcp.so](mcp.so) 已经注册了超过 15000 个 MCP 服务。
其目的是为了统一大模型与外部数据源和工具之间的通信协议，解决大模型数据孤岛的问题。

    简单来说 MCP 是扩展 AI 数据访问的一种简单的方式。

## MCP 的架构

MCP 的是 CS (Client - Server) 架构，这样一个 `主机(host)` 可以链接多个 `服务(Servers)`。

以下是 MCP 协议的架构组成概念:

| Concept          | 名称       | 解释 |
| :--------------: | :--------: | :--- |
| MCP Hosts        | MCP 主机   | 希望可以通过 MCP 访问数据的应用，比如 ClaudeDesktop, IDE, Cursor 等 |
| MCP Clients      | MCP 客户端  | 与服务器保持 1:1 链接的协议客户端 |
| MCP Servers      | MCP 服务端  | 轻量级程序，每个程序都通过标准化的模型上下文协议公开特定功能 |
| Local Data Source| 本地数据    | MCP 可以安全访问的文件、数据库和服务 |
| Remote Services  | 远程服务    | MCP 可以链接的外部互联网服务。|


### 参考引用
- [Model Context Protocol](https://zhuanlan.zhihu.com/p/27327515233)
- [知乎：一文看懂：MCP（大模型上下文协议）](https://zhuanlan.zhihu.com/p/27327515233)
- [知乎：从0到1掌握 MCP：微软 AI 标准化协议学习指南深度解读](https://zhuanlan.zhihu.com/p/1914634098984615970)
- [Github：microsoft/mcp-for-beginners](https://github.com/microsoft/mcp-for-beginners)