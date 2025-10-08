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

### 开始

这一章，
我们要开始设计 PCB 的板材、定位板，
同时为了简化设计，
我决定使用三明治的堆叠方案。

这一章不涉及电路板的绘制。

### 绘制三明治类型的定位板

打开 [http://builder.swillkb.com/](http://builder.swillkb.com/), 
这里我们需要按照如下信息进行选择：
![上传配列信息](/learn/img/ln/key_board/diy_my_first_keyboard/03_pcb_create/pcb_create.png)

将我们从上篇文章中获取的 raw data 的内容直接粘贴到这个网站中。

1. Switch Type: PCB 轴体开孔类型，这里选择 `MX`.
2. Stabilizer Type: 卫星轴开口类型，建议选择 `Cherry Only`, 因为兼容 `Costar` 类型的开口会导致定位板无法卡主 Cherry 轴体，不太稳定。
3. Case Type: 因为不想把外壳设计的过于复杂，所以我这里直接考虑使用 Sandwich 类型。其他类型需要自己绘制、外壳模型。
4. USB Count: 因为不想搞蓝牙和电源，所以我这里仅考虑有线的类型，所以需要预留一个 USB 口。
5. Mount Holes： 我需要保留 4 个 2 毫米的螺丝孔。
6. Edge Padding: 边缘预留大小。

### References

#### 绘制键盘定位板的网站

[http://builder.swillkb.com/](http://builder.swillkb.com/)

