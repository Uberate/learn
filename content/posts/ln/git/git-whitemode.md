+++
date = '2025-08-17T22:16:00+08:00'
title = 'Git 的白名单模式'
tags = ["git", "Scripts"]
categories = ["技术"]
author = "Uberate"
+++

> 参考： [gitignore](https://rgbcu.be/blog/gitignore)

# 引发

如果你是使用 mac 的工作环境，
那么我相信你一定在你的代码仓库中提交过 `.DS_STORE` 这个文件。

或者你提交过 `.idea/xxx.xml` ?

总之，你总是会在你的代码中不小心带上了不应该带上的文件对吧。

如何解决呢？

# .gitignore 

使用如下 git 的 ignore 文件内容：
> .gitignore
```yaml
# ignore all file first
*

# allow pkg dir, for example go
!pkg/
!pkg/**
!pkg/**/*.go
!cmd
!cmd/**/
!cmd/**/*.go
!go.{mod,sum}
!*.md
```

这个 gitignore 会被转换为白名单模式，只有列在其中的文件可以被提交。