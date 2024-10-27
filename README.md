# GitHub Sentinel Go

GitHub Sentinel Go 是一款开源工具类 AI Agent，专为开发者和项目管理人员设计，能够定期自动获取并汇总订阅的 GitHub 仓库最新动态。其主要功能包括订阅管理、更新获取、通知系统、报告生成。

## 特性

- **订阅管理**：可以轻松地订阅和取消订阅 GitHub 仓库。
- **自动更新**：定期获取订阅仓库的最新动态（Pull Requests、Issues、Commits 等）。
- **事件查看**：通过 CLI 立即查看订阅仓库的最新活动。
- **使用 go-github 库**：依赖 GitHub 官方 API 客户端库 `go-github`。

## 安装与使用

1. 克隆仓库：`git clone https://github.com/zysuper/github-sentinel-go.git`
2. 安装依赖：`go mod tidy`
3. 构建可执行文件：
   - CLI 可执行文件：`go build -o github-sentinel-cli cmd/cli/main.go`
   - 定时任务可执行文件：`go build -o github-sentinel-scheduler cmd/scheduler/main.go`
4. 设置 GitHub 和 OpenAI API 的访问令牌：
   ```bash
   export GITHUB_TOKEN=your_github_token_here
   export OPENAI_API_KEY=your_openai_api_key_here
   ```
5. 启动 REPL 模式的 CLI：`./github-sentinel-cli`
6. 启动定时任务模式：
   ```bash
   ./github-sentinel-scheduler -days 7 -interval 24h
   ```

#### 运行 CLI 模式

编译完成后，您可以通过以下命令来运行工具并进入交互式 REPL 环境：

```bash
./github-sentinel
```

在 REPL 模式下，您可以执行以下命令：

- `subscribe owner/repo`：订阅 GitHub 仓库。
- `unsubscribe owner/repo`：取消订阅 GitHub 仓库。
- `list`：列出所有已订阅的仓库。
- `update`：手动获取所有订阅仓库的更新。
- `help`：显示可用命令列表。
- `exit`：退出 REPL 环境。

##### 示例

```
> subscribe octocat/Hello-World
Successfully subscribed to: octocat/Hello-World

> list
Subscribed repositories:
- octocat/Hello-World

> update
Fetching updates...
```

#### 定时任务模式：

`./github-sentinel-scheduler -days 30 -interval 1h`

抓取过去 30 天的更新，每小时自动更新一次。



### 配置环境变量

在开始使用 GitHub Sentinel Go 之前，请确保已设置 GitHub 访问令牌。

您可以在 GitHub [Personal Access Tokens](https://github.com/settings/tokens) 页面创建一个新的令牌，并将其存储为环境变量。令牌需要具备读取仓库的权限。
