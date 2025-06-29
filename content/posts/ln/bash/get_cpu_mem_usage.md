+++
date = '2025-06-29T23:08:43+08:00'
draft = false 
title = 'Linux 中获取 CPU 和内存的使用情况'
categories = ["技术"]
tags=["bash", "Scripts"]
+++

```bash
# CPU 利用率
ps -p ${PID} -o %cpu
# 内存
ps -p ${PID} -o rss
```
