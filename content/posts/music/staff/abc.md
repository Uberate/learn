+++
date = '2025-06-07T01:28:17+08:00'
title = 'Abc 记谱法'
categories=['音乐']
tags = ['abc']
+++

## 头信息

```abc
X: 文件标号
T: 标题
M: 节拍 2/4
L: 谱号编辑音长单位 1/8
C: 作者
N: 备注 Ub(s)erate ABC
K: 调号 C
```

## ABC 内容
{{< abc >}}
M: 2/4
L: 1/4
K: C
CDEF|GABc|defg|abc'd'|e'f'g'a'|b'
{{< /abc >}}

- 延音: `{音符 音长}>`
- 连音: `(start end)`

## ABC 高低音
{{< abc >}}
M: 2/4
L: 1/4
K: C
V:1 Program 1 0 %piano
V:2 Program 1 0 octave=-2 bass $piano
V:1
CDEF|GABc|defg|abc'd'|e'f'g'a'|b'
V:2
CDEF|GABc|defg|abc'd'|e'f'g'a'|b'
{{< /abc >}}