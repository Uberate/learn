+++
date = '2025-06-15T18:57:15+08:00'
title = '读 《Boredom Over Beauty: Why Code Quaility is Code Security》'
categories = ["阅读"]
tags = ["短篇阅读", "代码", "代码质量", "代码安全"]
+++

> 原文 [Boredom Over Beauty: Why Code Quality is Code Security](https://blog.asymmetric.re/boredom-over-beauty-why-code-quality-is-code-security/#the-pragmatism-of-prevention) -- John Saigle

{{< lead >}}
代码质量就是代码安全
{{< /lead >}}

文中作者阐述了 `代码安全` 与 `代码质量` 的关联，作者认为两者是密不可分的。

### 优秀的、可维护的代码往往会不容易引入安全问题

管理混乱的代码由于其可维护性比较低，容易遭受安全问题的注入。
此外，没有结构化和一致的代码风格会导致审核人员无法有效、及时的发现问题。

由此可以得出一个结论：
1. 代码需要保持**风格统一**。

### 简单的代码一般出现的问题比较少

> The most secure code isn't beautiful, clever, or elegant-it's boring. Go, a
> programming language that exemplifies this philosophy, intentionally omits
> features that might enable expressive or elegant solutions in favor of simplicity
> and readability. This apparent limitation produces code that's easier to audit,
> harder to misuse, and less prone to security vulnerabilities.

翻译

> 最安全的代码不是美观、聪明或优雅的，而是无聊。Go 是一种体现这一理念的编程语
> 言，他故意省略了可能实现富有表现力或优雅的解决方案的功能，以支持简单性和可
> 读性。这种明显的限制会产生更容易审计、更难误用且不易出现安全漏洞的代码。

读到这里，突然想起来 `Golang` 社区彻底放飞自我，停止对 `if err != nil {} `的优化，[[ On | No ] syntactic support for error handling](https://go.dev/blog/error-syntax)。

实打实的来说，从我目前的开发经验来看，越是上层的代码就越需要处理错误。如果放任不管一定是不行的。
这也是 Golang 被人诟病的一个地方，err 总是作为一个返回值来表述。
但是，大部分的社区库对 err == nil 的行为是UB（未定义行为）的，所以很容易引发下面这个写法。

```golang
func main() {
    anObjPtr, err := third.Get()
    if err != nil {
        os.Exit(1) // some error here.
    }

    // 但是 test 也可以是nil 。所以一旦这种情况发生，就会导致 panic
    anObjPtr.SetName() // panic: anObjPtr == nil.
}
```

话题扯远了。让我再回过头来看，作者的本意就是让错误暴露出来，必须要处理，而不是什么地方通过 `美丽` 的写法隐藏起来。

不只是错误的处理，所有代码必须要做到单一职责。
我记得我的领导对优秀代码的看法：

    优秀的代码一定可以很轻松的写出它的单元测试。

我非常认同。