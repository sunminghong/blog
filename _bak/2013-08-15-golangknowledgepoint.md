---
layout: post
categories: 
- golang
- work
title: golang知识点笔记一
description: 记录一些学用go 过程中的一些变态的、难以遇到的点滴。
keywords: gob,json,serialize,序列化,编码,解码,golang
---

###前言

###测试机配置


###测试代码

**serialize.go**
{% highlight go lineno%}
{% endhighlight %}

**serialize_b_test.go**
{% highlight go lineno%}
{% endhighlight %}

1. 玩家信息数据，包括玩家身份数据、玩家全局数值类数据，如玩家id，角色名，元宝等
2. 主要游戏数据，包括玩家现在状态、成长、成就、道具等数据，是玩家游戏的结晶
3. 日志类数据，包括一般日志（如战斗日志、获奖日志、登陆日志等）和核心日志（如充值记录、消费记录、领奖记录等）