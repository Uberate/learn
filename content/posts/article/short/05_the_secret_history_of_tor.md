+++
date = '2025-08-17T20:26:00+08:00'
title = '读 《The Secret History of Tor》'
categories = ["阅读"]
tags = ["短篇阅读", "安全", "Tor"]
+++

> 原文 [The Secret Hisory of Tor](https://thereader.mitpress.mit.edu/the-secret-history-of-tor-how-a-military-project-became-a-lifeline-for-privacy/) -- Ben Collier

# 警示

这篇文章对于技术的探究没有很深，
不过他让我有了一些额外的思考。

Tor 始于 NRL(美国海军研究实验室)。
其目标是通过 Tor 建立一个安全的通信方式。

所谓安全，
是指不会被跟踪，
数据被补货后也不会泄露太多的数据。
基于以上两点，
tor 出现了。
其工作方式依赖洋葱网络。

不过，
他的工作有一个前提，存在大量的用户基础。
如果用户很少，那么可以被用来做信息传输的中间人就越少。

所以，
我认为美国为了建立更好的间谍体系，
其应该会大力宣传 Tor。

越多的用户意味着越安全。
我深以为然。

{{< lead >}}
There are other implications,
as well.
For a CIA agent to use Tor without suspicion in non-U.S. nations,
for example, 
there would need to be plenty of citizens in these nations using Tor for everyday internet browsing.
Similarly,
if the only users in a particular country are whistleblowers, 
civil rights activists and protesters, 
the government may well simply arrest anyone connection to your anonnymity network.
As a result,
an onion routing system had to be open to as wide a range of users and maintainers as possible,
so that the mere fact that someone was using the system wouldn't reveal anything about their idenity or their affiliations.

这也带来了其他影响。
例如，
中央情报局（CIA）的特工若想在非美国国家使用“洋葱路由”（Tor）而不被怀疑，
这些国家就需要有大量公民日常使用Tor进行网络浏览。
同样，
如果某个国家使用Tor的只有举报人、
民权活动人士和抗议者，
该国政府很可能会直接逮捕任何连接到这个匿名网络的人。
因此，
洋葱路由系统必须尽可能向广泛的用户和维护者开放，
这样一来，
仅从某人使用该系统这一事实，
不会暴露其身份或所属关系。
{{< /lead >}}

# 工作原理

`onion` 网络，
目标信息被加密在三层网络上。

其实际转发的地址也并非目的地。
通过绕远的方式来建立连接。

同时将数据包拆成多个，
以此保证单数据包被拦截后，
其数据仍是安全的。

{{< lead >}}
The routing information used to navigate the internet is first hidden under three layers of encryption, like a Russian doll.

用于浏览互联网的路由信息首先像俄罗斯套娃一样，
被隐藏在三层加密之下。
{{< /lead >}}

# 赛博精神

Cypherpunks warned that the internet cloud quickly turn from a utopian dream into an authoritarian nightmare.

事实上，中国网络状态已经是全透明的状态了。
一方面，这建立了全世界最为安全的社会体系。
同时，这也完全牺牲了个人隐私。