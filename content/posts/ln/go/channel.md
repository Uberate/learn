+++
date = '2025-06-02T19:10:41+08:00'
draft = true
title = 'Channel'
tags = ["ln", "go"]
categories = ["技术"]
author = "Uberate"
topic = "深入 Go"
+++

# Golang Channel

## ReadNil

```go
package main

func main() {
    var nilChan chan struct{} = nil
    <- nilChan
}
```