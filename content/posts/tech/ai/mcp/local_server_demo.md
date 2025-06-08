+++
date = '2025-06-04T22:39:13+08:00'
title = '学习 MCP - 本地 Server Demo'
tags = ['AI', 'MCP']
categories = ['技术']
+++

这篇文章将带你创建一个简单的 MCP Local Server 服务。

代码已经放到我的 Github 仓库里面了：

{{< github repo="Uberate/lear-mcp" >}}

> 需要注意，不同的操作系统的准备工作是不同的，详细的信息可以访问 [https://modelcontextprotocol.io/quickstart/server](https://modelcontextprotocol.io/quickstart/server) 了解
> 我的环境是 `MAC` 的操作系统。

## 环境准备

### 1. 安装 uv

执行以下指令：

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh

# installing to /Users/xxx/.local/bin
#   uv
#   uvx
# everything's installed!
#
# To add $HOME/.local/bin to your PATH, either restart your shell or run:
#
#     source $HOME/.local/bin/env (sh, bash, zsh)
#     source $HOME/.local/bin/env.fish (fish)

# 让当前的 Context 生效
source $HOME/.local/bin/env 
```

将 `$HOME/.local/bin` 添加到环境路径 `PATH` 中。笔者用的是 `ZSH` 所以直接添加到 `~/.zshrc`

```bash
# 备份
cp ~/.zshrc ~/.zshrc_bk

# 添加
cat << EOF >> ~/.zshrc

# uv tool for mcp
PATH="$HOME/.local/bin:$PATH"
EOF
```



### 2. 创建一个文件夹，用来存放我们的代码

```bash
# 为我们的 Demo 项目创建一个目录
uv init mcp_local_server_demo
cd mcp_local_server_demo

# 创建虚拟环境并使用该虚拟环境
uv venv
source .venv/bin/activate

# 安装 mcp 依赖
uv add "mcp[cli]"

# 创建我们的服务代码文件
touch local_server.py
```

![项目目录](/learn/img/tech/ai/mcp/local_server_demo/project_layout.png "项目目录")

---

### 3. 编写服务

通过 `步骤2` 我们已经创建了 `local_server.py` 代码文件。
现在我们需要实现这个代码。

### 4. 实现目标
我们来实现一个基于 `ping` 命令的网络延时检测服务，用于帮助用户快速检查网络状态。

    当然，用户不太可能在没有网络的情况下使用 MCP？

为 `local_server.py` 添加如下代码:

```py
from mcp.server.fastmcp import FastMCP
import socket
import time
from concurrent.futures import ThreadPoolExecutor

# Initialize FastMCP server
mcp = FastMCP("local_server")

site_list = [
    "www.baidu.com",
    "www.google.com",
    "www.github.com", 
    "www.bilibili.com"]

@mcp.tool()
def pp() -> str:
    def format_result(site):
        return site + "\t" + tcp_ping(site) + "\n"

    with ThreadPoolExecutor(max_workers=4) as executor:
        results = list(executor.map(format_result, site_list))
    return "".join(results)

def tcp_ping(host: str, port: int = 80, timeout: float = 3.0) -> str:
    try:
        start = time.time()
        sock = socket.create_connection((host, port), timeout=timeout)
        sock.close()
        delay = (time.time() - start) * 1000  # 转为毫秒
        return f"{delay:.2f}ms"
    except Exception as e:
        return str(e)

if __name__ == "__main__":
    # Initialize and run the server
    mcp.run(transport='stdio')
```

### 5. 测试

我们可以借助[MCP 分析工具](https://github.com/modelcontextprotocol/inspector)进行测试

```bash
npx @modelcontextprotocol/inspector mcp run local_server.py
```

他会返回一个代理地址:
![UI](/learn/img/tech/ai/local_server_demo/open_inspect.png "WEB UI Inspect")

在浏览器打开后，按照如下顺序点击:

![WEB_INSPECT_CONNECT](/learn/img/tech/ai/local_server_demo/web.png)

可以看到我们的 TOOL 已经可以正常使用了。

### 6. 配置 Cursor

笔者这里使用的是 `cursor` 所以这里把 MCP 配置到 Cursor 进行一次验证。

方式比较简单，在项目的根目录下执行一下指令:
```bash
# repo root
mkdir .cursor
cat << EOF > .cursor/mcp.json
{
    "mcpServers": {
        "local_server": {
            "command": "./mcp_local_server_demo/.venv/bin/mcp",
            "args": [
                "run",
                "./mcp_local_server_demo/local_server.py"
            ]
        }
    }
}
EOF
```
> 这里配置的项目级别的 MCP。

然后打开 cursor 的配置页面:
![Cursor config](/learn/img/tech/ai/local_server_demo/cursor_config.png)
可以看到这个 MCP 提供了一个名为 `pp` 的 tools。

打开 `cursor` 的对话框，输入：
```bash
# 输入
用PP 检查一下我的网络状态。

# Cursor 响应：
根据PP工具的检测结果，您的网络状态如下：
✅ www.baidu.com - 153.44ms（正常）
❌ www.google.com - 超时（无法访问）
✅ www.github.com - 235.94ms（正常，但响应稍慢）
✅ www.bilibili.com - 147.43ms（正常）
总结：
您的网络连接基本正常
国内网站（百度、B站）访问速度良好
Google网站无法访问，可能是网络环境限制
GitHub可以访问但响应时间较长
网络整体状况良好，只是部分海外网站访问受限。
```

![](/learn/img/tech/ai/local_server_demo/cursor_call_result.png)


# 参考引用
- [https://modelcontextprotocol.io/quickstart/server](https://modelcontextprotocol.io/quickstart/server)