+++
date = '2025-08-31T22:03:07+08:00'
draft = false
title = 'Claude code router'
categories = ['技术']
tags = ['ClaudeCode', 'GitRepos', "AI"]
+++

## 背景

由于 ClaudeCode 的一系列原因，
可能无法快速在这个地区合理使用。
所以找到了一个可以路由到其他模型的方法。

### GitRepo

{{< github repo="musistudio/claude-code-router" >}}

## 安装

claud-code-router 主要解决模型问题，
所以在开始前你必须已经安装了 `claude-code`。

```bash
npm install -g @anthropic-ai/claude-code
```

然后你还需要安装路由器
```bash
npm install -g @musistudio/claude-code-router
```

### 使用

你可以通过 **C**laude **C**ode **R**outer 命令来启动 UI 界面。

```bash
ccr ui
```

![UI Demo](/learn/img/tech/ai/claude_code_router/1.png)

根据 UI 的提示填写相关信息即可。

我选择的是火山提供商的 DeepSeek。

SK 是方舟模型中的 API Key。

然后，启动：
```bash
ccr start
```

然后在你的项目启动 CaludeCode
```bash
ccr code
```