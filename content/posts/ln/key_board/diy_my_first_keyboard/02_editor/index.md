+++
date = '2025-10-05T13:50:00+08:00'
draft = true
title = 'DIY 键盘 02: 准备工作以及设计配列'
categories = ['技术']
tags=['学习', '键盘', 'HHKB']
series = ["DIY 键盘"]
series_order = 2
+++


你可以在我的仓库中找到我所使用的全部文件：
{{< github repo="Uberate/diy_hhkb" >}}

### References

- [keyboard-layout-editor.com](https://keyboard-layout-editor.com/#/) 是一个编辑键盘配列的工具网站。


### 开始

首先，
我们需要学习什么是配列，
当然，
如果你比较急或者对这些知识毫无兴趣，
你可以打开我的 `Github` 直接获取 `HHKB` 的键盘配列信息。

我这里使用的是 [keyboard-layout-editor.com](https://keyboard-layout-editor.com/#/) 进行的配列编辑，
你可以直接上传[我的](https://github.com/Uberate/diy_hhkb/blob/main/keyboard-layout-editor.com/hhkb.json) `hhkb.json`
文件来获取我的 HHKB 配列信息，
然后进行二次改造等。

或者你已有 keyboard-layout-editor.com 格式的 HHKB 配列格式的话，
你可以跳过这篇文章。

![上传配列信息](/learn/img/ln/key_board/diy_my_first_keyboard/02_editor/upload_json_file.png)

### 什么是配列

首先是关于**配列**这个单词。

一开始我看这个词的时候看得我一头雾水，
然后我猜测这个应该是音译过来的，
但是我想了一圈也没有想到这个对应的英语单词。
后来查找了一些资料，
才知道配列是日语音译过来的，
对应的英文单词是 `array`（🤔），
数组的意思。

实际上，
在 DIY 键盘的社区中，
配列一词与数组感念无关，
实际上是指的**布局**（Layout），
或者说**排列**更加合适。
我的确没有找到为什么不使用**排列**这个单词的原因。
不过社区已经形成概念，
理解就好。

> 为了保证后续的一直性，如果没有特殊的必要性，我会使用配列代替键盘布局一词，避免概念复杂化。

#### 键盘的区域

在介绍常见的配列之前，
我们先看一下一个键盘包括什么区域。
毕竟没有人喜欢将按键的位置调整或删除，
但是区域可以。
比如一些 `87 配列` 的键盘是没有 `小键盘（数组按键）` 区域的。

1. 主键盘区（字母数字区）: 核心输入区，负责文字、数字、符号的基础输入 26 个字母键、10 个数字键（顶排）、Shift/Ctrl/Alt
   等修饰键、空格键、回车键
2. 功能键区（F 区）: 快捷操作触发，可单独或配合修饰键实现功能 F1-F12（部分全尺寸键盘含 F13-F24），默认功能如 F5 刷新、F11 全屏
3. 编辑控制区: 文本 / 光标控制，减少鼠标依赖 Insert/Delete（插入 / 删除）、Home/End（行首 / 行尾）、Page Up/Page
   Down（翻页）、方向键（↑↓←→）
4. 数字小键盘区（ numpad ）: 高效数字输入，含基础运算功能 0-9 数字键、+/-×÷ 运算键、Enter（数字回车）、Num Lock（锁定键）
5. 状态指示灯区: 显示键盘工作状态，直观反馈功能开关 Caps Lock（大小写锁定）、Num Lock（小键盘锁定）、Scroll Lock（滚动锁定）

*如下图所示*

![104 配列](/learn/img/ln/key_board/diy_my_first_keyboard/02_editor/104_key_board_layout.png)

#### 常见配列

这里我整理了一下常见的配列表：
| 配列类型 | 键位数量 | 核心特征 |
|:---------------:|:---------:|:------------------------------------------|
| 104/108 配列（全尺寸） | 104-108 键 | 保留字母区、数字小键盘、编辑区、F 区，功能最完整 |
| 98 配列（96%） | 96-98 键 | 缩减编辑区冗余按键，保留小键盘，兼顾便携与数字输入 |
| 87 配列（TKL） | 87 键 | 移除小键盘，保留 F 区与编辑区，桌面占用率降低 20% |
| 75 配列（70%-75%） | 82-84 键 | 合并编辑区与方向键，部分型号增加音量旋钮等快捷控件 |
| 65 配列（65%） | 67-68 键 | 移除 F 区，压缩编辑区，保留独立方向键 |
| 60 配列（60%） | 61 键 | 仅保留字母区，依赖 FN 组合键实现功能扩展 极简桌面、便携外出 |
| 40 配列（40%） | 40-45 键 | 移除数字区与大部分符号键，极致紧凑 极端便携需求（如户外直播） 客制化 40% 套件 |

### 创建一个 HHKB 配列

当然，
你总是可以在互联网上找到你需要的东西。
如果没有找到，
你可以用我的。
我这里使用的是 HHKB 标准的配列。
我发现目前社区中存在大量 60 配列的来当做 HHKB 配列。
然后社区中叫 HHKB 标准配列为 TrueHHKB 配列。
注意甄别你需要的类型。

这里注意保存 `raw data` 和 `json` 文件。

区别方法是看空格长度，
理论上 `6U` 空格不会过多的侵占右侧 `CTRL` ，
标准的 `HHKB` 配列中，
`CTRL` 在右侧 `Shift` 的左面下发，
而 60 配列中，
由于空格是 `7U` 的，
所以 `CTRL` 几乎到了右侧 `Shift` 的中间甚至靠右侧的下方。

![上传配列信息](/learn/img/ln/key_board/diy_my_first_keyboard/02_editor/upload_json_file.png)

[hhkb.json](https://github.com/Uberate/diy_hhkb/blob/main/keyboard-layout-editor.com/hhkb.json)
{{< codeimporter url="https://raw.githubusercontent.com/Uberate/diy_hhkb/refs/heads/main/keyboard-layout-editor.com/hhkb.json" type="json" >}}