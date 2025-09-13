+++
date = '2025-09-14T00:11:56+08:00'
title = 'Use singular nouns for database table names'
categories = ["阅读"]
tags = ["短篇阅读","数据库"]
+++

> 原文 [Use singular nouns for database table names](https://www.teamten.com/lawrence/programming/use-singular-nouns-for-database-table-names.html) -- Lawrence Kesteloot

按照现代的数据库建设来说，
大多数的规范，
甚至包括一些基础库（Gorm等）都会直接将表名命名为复数形式。

我对作者的观点持反对意见，着重对第二部分，单数形式的必要性和理由进行阐述：

### 1. 命名的就是表名，而不是关系 
作者认为，
实际上不是在给表命名，
而是在给关系命名。

对于这个问题，
我是这样思考的：

毋庸置疑的一点是，
我们就是在给表 `table` 命名，
不是关系 `relation?`。
*（作者认为是在给关系命名，
我无法理解。）*

关系定义了我们如何组织表中的属性，
表的目的是记录某一类关系的具体数据，
而不是记录关系。

依旧是用户的这个例子，
看到这个表的时候，
我们的第一感觉是这个 `users` 里面有(系统)中所有的用户信息。
而如何看到 `user` 大多数的人都会觉得这是对 `users` 的关系抽象，
或者是具体某个用户，
总之，
`user` 无法代表用户集合。

至于 `user` 这个关系中包含哪些属性，
我认为作者说的没错，
这个是属性的关系，
这个关系不存在负数的场景。

### 2. 有些 SQL 会比较奇怪。

> 引用作者的 sql 语句

```sql
SELECT id, name
FROM user
JOIN country ON user.country_id = country.id
WHERE country.name = 'Canada';
```

从我个人的态度来看，
我觉得 users 反而是合理的，
理由：

`country_id` = `country.id` 的用户大概率并不是唯一的。
我相信这张 SQL 的目的应该是选择所有 `Canada` 国家的用户。
所以看起来 users 好像更加合理。

我思考了一下，
我认为作者可能是将 country_id 看做用户群体的属性感到难受。
从表意上来看，
这里的目的是当这个 `user` 的 `country_id` 为 `Canada` 就认为符合条件。
是以单个个体出发的。

我觉得各有各理。

### 负数关系的定义

这其实没有任何意义。

通过三范式解决后的结构关系，
应该不存在所谓负数的负数形式。
比如作者举的例子 `UserFacts`, 
如果从更合理的命名形式来看，
可以命名为 `UserFactLists` 。

显然，不会有人脱裤子放屁。


